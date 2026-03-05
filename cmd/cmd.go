package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Mwawaka/go-crazy/internal/config"
	"github.com/Mwawaka/go-crazy/ui"
	"google.golang.org/genai"
)

func Run() {
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

	terminalUI.Start()
	reader := bufio.NewReader(os.Stdin)

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

		output, err := makeRequest(ctx, formattedInput, client, c)

		if err != nil {
			terminalUI.PrintError(err)
			continue
		}

		terminalUI.PrintResponse(output, 30*time.Millisecond)
	}

}

func makeRequest(ctx context.Context, prompt string, client *genai.Client, c *config.Config) (string, error) {
	result, err := client.Models.GenerateContent(
		ctx,
		c.Model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			// MaxOutputTokens: 8000, // Limits how many tokens the model can return in a single response
		},
	)

	if err != nil {
		return "", fmt.Errorf("generating content: %w", err)
	}

	return result.Text(), nil
}
