package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Gemini models",
	Long:  "Fetch and display all available Gemini models from the API.",
	Run: func(cmd *cobra.Command, args []string) {
		// To implement
	},
}

func init() {
	modelsCmd.AddCommand(listCmd)
}
