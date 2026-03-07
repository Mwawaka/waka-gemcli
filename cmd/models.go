package cmd

import "github.com/spf13/cobra"

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage Gemini models",
	Long:  "List and select available Gemini models for use with GemCLI",
	Run: func(cmd *cobra.Command, args []string) {
		// To implement
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}
