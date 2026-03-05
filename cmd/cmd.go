package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Mwawaka/go-crazy/internal/config"
	"google.golang.org/genai"
)

func Run() {

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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ gemini ")
		input, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
			break
		}

		formattedInput := strings.TrimSpace(input)

		if formattedInput == "" {
			continue
		}

		output, err := makeRequest(ctx, formattedInput, client)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Fprintln(os.Stdout, output)
	}

}

func makeRequest(ctx context.Context, prompt string, client *genai.Client) (string, error) {
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			// MaxOutputTokens: 8000, // Limits how many tokens the model can return in a single response
		},
	)

	if err != nil {
		return "", fmt.Errorf("error generating content")
	}

	return result.Text(), nil
}

func GetUserInput() {

}
