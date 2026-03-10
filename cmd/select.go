package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a Gemini model to use",
	Long:  "Select a Gemini model to use for your chat sessions. The selected model will be saved to your configuration",
	Run: func(cmd *cobra.Command, args []string) {
		models, err := fetchModels()
		modelNames := make([]string, 0, len(models))

		if err != nil {
			terminalUI.PrintError(err)
			return
		}

		for _, m := range models {
			modelNames = append(modelNames, m.Name)
		}

		prompt := promptui.Select{
			Label: "Choose a model",
			Items: modelNames,
		}
		_, selectedModel, err := prompt.Run()

		if err != nil {
			terminalUI.PrintError(err)
			return
		}

		model = selectedModel

		if err := loader.Set("MODEL", selectedModel); err != nil {
			terminalUI.PrintError(err)
			return
		}

		terminalUI.PrintChunk(fmt.Sprintf("\n✅ Model set to %s\n", selectedModel))
	},
}

func init() {
	modelsCmd.AddCommand(selectCmd)
}
