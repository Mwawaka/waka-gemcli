package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/Mwawaka/waka-gemcli/internal/config"
	"github.com/Mwawaka/waka-gemcli/internal/ui"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var (
	// Used for flags
	model      string
	client     *genai.Client
	cfg        *config.Config
	ctx        context.Context
	terminalUI *ui.UI
	loader     *config.ViperLoader

	rootCmd = &cobra.Command{
		Use:   "gemcli",
		Short: "A cyberpunk themed Gemini CLI client",
		Long:  "GemCLI is a cyberpunk styled command line interface for interacting with Google Gemini directly from your terminal. No browser. No UI. Just you, the terminal, and the AI.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			terminalUI = ui.NewUI(os.Stdin, os.Stdout, os.Stderr)
			loader, err = config.NewViperLoader()

			if err != nil {
				return err
			}

			cfg, err = loader.Load()
			ctx = context.Background()

			if err != nil {
				// Checks whether it is a first run
				if errors.Is(err, config.ErrConfigNotFound) {
					cfg, err = terminalUI.PromptInitialConfig()

					if err != nil {
						return err
					}

					if err = loader.Set("GOOGLE_API_KEY", cfg.APIKey); err != nil {
						return err
					}

					if err = loader.Set("MODEL", cfg.Model); err != nil {
						return err
					}

				} else {
					return err
				}
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
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "gemini-2.5-flash", "Default Gemini model to use")
}

// Cobra prints errors to stderr by default
