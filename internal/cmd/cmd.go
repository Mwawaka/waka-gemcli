package cmd

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Mwawaka/go-crazy/internal/config"
	"github.com/Mwawaka/go-crazy/internal/ui"

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

		output, err := makeRequest(ctx, formattedInput, chat)

		if err != nil {
			terminalUI.PrintError(err)
			continue
		}

		terminalUI.PrintResponse(output, 30*time.Millisecond)
	}

}

func makeRequest(ctx context.Context, prompt string, chat *genai.Chat) (string, error) {
	// result, err := chat.SendMessage(ctx, genai.Part{
	// 	Text: prompt,
	// })

	// if err != nil {
	// 	return "", fmt.Errorf("generating content: %w", err)
	// }

	// return result.Text(), nil
	return "test", nil
}
