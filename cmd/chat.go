package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session with Gemini",
	Long:  "Start a cyberpunk styled  interactive chat session with Gemini. Maintains conversation history across messages. Type !help for available commands",
	Run: func(cmd *cobra.Command, args []string) {
		if model == "" {
			model = cfg.Model
		}

		chat, err := client.Chats.Create(
			ctx,
			model,
			&genai.GenerateContentConfig{
				//
			},
			nil,
		)

		if err != nil {
			terminalUI.PrintError(err)
			return
		}

		// terminalUI.Start()
		reader := bufio.NewReader(os.Stdin)

	loop:
		for {

			terminalUI.PrintPrompt()
			input, err := reader.ReadString('\n')

			if err != nil {
				if err == io.EOF {
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
					break loop
				case "!clear":
					newChat, err := client.Chats.Create(
						ctx,
						model,
						&genai.GenerateContentConfig{},
						nil,
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

			if err = makeRequest(ctx, formattedInput, chat, terminalUI.PrintChunk); err != nil {
				terminalUI.PrintError(err)
				continue
			}

			terminalUI.PrintChunk("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

func makeRequest(ctx context.Context, prompt string, chat *genai.Chat, onChunk func(string)) error {
	var apiErr genai.APIError
	chatIterator := chat.SendMessageStream(ctx, genai.Part{
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
