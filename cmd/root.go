package cmd

import (
	"context"
	"os"

	"github.com/Mwawaka/waka-gemcli/internal/config"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var (
	// Used for flags
	model  string
	client *genai.Client
	cfg    *config.Config
	ctx    context.Context

	rootCmd = &cobra.Command{
		Use:   "gemcli",
		Short: "A cyberpunk themed Gemini CLI client",
		Long:  "GemCLI is a cyberpunk sytled command line interface for interacting with Google Gemini directly from your terminal. No browser. No UI. Just you, the terminal, and the AI.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			loader := &config.ViperLoader{
				FileType: ".env",
			}
			cfg, err = loader.Load()
			ctx = context.Background()

			if err != nil {
				return err
			}

			client, err = genai.NewClient(ctx, &genai.ClientConfig{
				APIKey:  cfg.APIKey,
				Backend: genai.BackendGeminiAPI,
			})

			if err != nil {
				return err
			}

			return nil
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&model, "model", "gemini-2.5-flash", "Default Gemini model to use")
}

// Cobra prints errors to stderr by default
