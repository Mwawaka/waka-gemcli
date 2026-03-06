package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Mwawaka/waka-gemcli/internal/config"
	"github.com/Mwawaka/waka-gemcli/internal/ui"
	"google.golang.org/genai"
)

func Run() {
	var chat *genai.Chat
	terminalUI := ui.NewUI(os.Stdout, os.Stderr)
	cfg := &config.ViperLoader{
		FileType: ".env",
	}
	c, err := cfg.Load()
	ctx := context.Background()

	if err != nil {
		log.Fatalln(err)
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  c.APIKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatal(err)
	}

	chat, err = client.Chats.Create(
		ctx,
		c.Model,
		&genai.GenerateContentConfig{
			// MaxOutputTokens: 8000, // Limits how many tokens the model can return in a single response
		},
		nil, // should pass history later
	)

	if err != nil {
		log.Fatal(err)
	}

	terminalUI.Start()
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
					c.Model,
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

		err = makeRequest(ctx, formattedInput, chat, terminalUI.PrintChunk)

		if err != nil {
			terminalUI.PrintError(err)
			continue
		}

		terminalUI.PrintChunk("\n")

	}

}

func makeRequest(ctx context.Context, prompt string, chat *genai.Chat, onChunk func(string)) error {
	chatIter := chat.SendMessageStream(ctx, genai.Part{
		Text: prompt,
	})

	for result, err := range chatIter {
		if err != nil {
			return fmt.Errorf("generating content: %w", err)
		}

		onChunk(result.Text())
	}

	return nil

}


