package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/genai"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage Gemini models",
	Long:  "List and select available Gemini models for use with GemCLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}

func fetchModels() ([]*genai.Model, error) {
	var models []*genai.Model
	var apiErr genai.APIError

	for model, err := range client.Models.All(ctx) {
		if err != nil {
			if errors.As(err, &apiErr) {
				return nil, fmt.Errorf("\n[%d %s] request failed - %s\n Run gemcli config set GOOGLE_API_KEY [API key] to configure your API key\n", apiErr.Code, apiErr.Status, apiErr.Message)
			}
			return nil, err
		}

		models = append(models, model)
	}

	return models, nil
}
