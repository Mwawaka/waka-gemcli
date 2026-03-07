package cmd

import "github.com/spf13/cobra"

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a Gemini model to use",
	Long:  "Select a Gemini model to use for your chat sessions. The selected model will be saved to your configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// To implement
	},
}

func init() {
	modelsCmd.AddCommand(selectCmd)
}
