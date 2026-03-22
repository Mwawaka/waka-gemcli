package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"iter"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Mwawaka/waka-gemcli/internal/history"
	"github.com/Mwawaka/waka-gemcli/internal/ui"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

type ChatSession interface {
	SendMessageStream(ctx context.Context, parts ...genai.Part) iter.Seq2[*genai.GenerateContentResponse, error]
}

var (
	resume bool

	chatCmd = &cobra.Command{
		Use:   "chat",
		Short: "Start an interactive chat session with Gemini",
		Long:  "Start a cyberpunk styled  interactive chat session with Gemini. Maintains conversation history across messages. Type !help for available commands",
		Run: func(cmd *cobra.Command, args []string) {
			var hist []*genai.Content
			var prompt string

			if model == "" {
				model = cfg.Model
			}

			hfg, err := history.NewHistory()

			if err != nil {
				terminalUI.PrintError(err)
				return
			}

			if resume {
				hist, err = hfg.LoadHistory()

				if err != nil && !errors.Is(err, fs.ErrNotExist) {
					terminalUI.PrintError(err)
					return
				}

				hist = nil
			}

			chat, err := client.Chats.Create(
				ctx,
				model,
				&genai.GenerateContentConfig{
					//
				},
				hist,
			)

			if err != nil {
				terminalUI.PrintError(err)
				return
			}

			setupSignalHandler(hfg, chat, terminalUI)

			// terminalUI.Start()
			reader := bufio.NewReader(os.Stdin)

			// Detecting whether input is piped or from a character device(keyboard)
			fileInfo, err := os.Stdin.Stat()
			if err != nil {
				terminalUI.PrintError(err)
				return
			}

			// 0 -piped, 1- character device
			if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
				content, err := io.ReadAll(os.Stdin)

				if err != nil {
					terminalUI.PrintError(err)
					return
				}

				if len(args) == 0 {
					terminalUI.PrintError(fmt.Errorf("please provide a question when using piped content"))
					return
				}

				prompt = fmt.Sprintf("Content:\n%s\n\nQuestion: %s", content, args[0])

				if err := makeRequest(ctx, chat, prompt, terminalUI.PrintChunk); err != nil {
					terminalUI.PrintError(err)
					return
				}

				// Solving: pipe and keyboard  reading from os.Stdin
				tty, err := os.Open("/dev/tty")

				if err != nil {
					terminalUI.PrintError(err)
					return
				}

				defer tty.Close()
				reader = bufio.NewReader(tty)
			}

		loop:
			for {
				terminalUI.PrintPrompt()
				input, err := reader.ReadString('\n')

				if err != nil {
					if err == io.EOF {
						if resume {
							if err := hfg.SaveHistory(chat.History(false)); err != nil {
								terminalUI.PrintError(err)
							}
						}

						break
					}

					terminalUI.PrintError(err)
					break
				}

				formattedInput := strings.TrimSpace(input)

				if formattedInput == "" {
					continue
				}

				if strings.HasPrefix(formattedInput, "!") {
					switch formattedInput {
					case "!exit":
						// save comprehensive history
						if err := hfg.SaveHistory(chat.History(false)); err != nil {
							terminalUI.PrintError(err)
						}
						break loop
					case "!clear":
						newChat, err := client.Chats.Create(
							ctx,
							model,
							&genai.GenerateContentConfig{},
							nil, //fresh start - no history
						)

						if err != nil {
							terminalUI.PrintError(err)
							continue
						}

						chat = newChat
					case "!help":
						terminalUI.PrintHelp()
					default:
						terminalUI.PrintInvalidInput(formattedInput)
					}

					continue
				}

				terminalUI.PrintChunk("\n🧠 GemCLI says:\n\n")

				if err = makeRequest(ctx, chat, formattedInput, terminalUI.PrintChunk); err != nil {
					terminalUI.PrintError(err)
					continue
				}

				terminalUI.PrintChunk("\n")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().BoolVarP(&resume, "resume", "r", false, "Resumes previous chat session")
}

func makeRequest(ctx context.Context, chatSession ChatSession, prompt string, onChunk func(string)) error {
	var apiErr genai.APIError
	chatIterator := chatSession.SendMessageStream(ctx, genai.Part{
		Text: prompt,
	})

	for result, err := range chatIterator {
		if err != nil {
			if errors.As(err, &apiErr) {
				return fmt.Errorf("\n[%d %s] request failed - check your quota and retry shortly\n", apiErr.Code, apiErr.Status)
			}

			return fmt.Errorf("generating content: %w", err)
		}

		onChunk(result.Text())
	}

	return nil
}

// setupSignalHandler handles saving history when the user presses CTRL + C
func setupSignalHandler(hfg *history.History, chat *genai.Chat, ui *ui.UI) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan

		if err := hfg.SaveHistory(chat.History(false)); err != nil {
			ui.PrintError(err)
		}

		os.Exit(0)
	}()
}
