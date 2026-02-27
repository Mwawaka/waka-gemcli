package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Mwawaka/go-crazy/internal/config"
	"google.golang.org/genai"
)

func main() {

	cfg, err := config.GetConfig(&config.ViperLoader{
		FileType: ".env",
	})
	ctx := context.Background()

	if err != nil {
		log.Fatalln(err)
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.APIKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text("Can I build CLI tools using golang"),
		&genai.GenerateContentConfig{
			// MaxOutputTokens: 8000, // Limits how many tokens the model can return in a single response
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, result.Text())
}
