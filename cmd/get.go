package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			terminalUI.PrintError(fmt.Errorf("invalid number of arguments"))
			return
		}

		key, err := loader.Get(args[0])

		if err != nil {
			terminalUI.PrintError(err)
			return
		}

		terminalUI.PrintChunk(key)
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
