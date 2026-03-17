package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Long:  "Sets a configuration key-value pair in ~/.cofig/gemcli/config.yaml. Example: gemcli config set GOOGLE_API_KEY mykey",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			terminalUI.PrintError(fmt.Errorf("invalid number of arguments"))
			return
		}

		if err := loader.Set(args[0], args[1]); err != nil {
			terminalUI.PrintError(err)
			return
		}

		terminalUI.PrintChunk(fmt.Sprintf("\n✅ %s updated successfully\n", args[0]))
	},
}

func init() {
	configCmd.AddCommand(setCmd)
}
