package cmd

import (
	"os"

	"github.com/Mwawaka/waka-gemcli/internal/ui"
	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Gemini models",
	Long:  "Fetch and display all available Gemini models from the API.",
	Run: func(cmd *cobra.Command, args []string) {
		var models []*genai.Model
		terminalUI := ui.NewUI(os.Stdout, os.Stderr)

		for model, err := range client.Models.All(ctx) {
			if err != nil {
				terminalUI.PrintError(err)
				return
			}
			models = append(models, model)
		}

		terminalUI.PrintModelInfo(models)
	},
}

func init() {
	modelsCmd.AddCommand(listCmd)
}
