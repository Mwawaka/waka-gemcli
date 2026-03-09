package cmd

import "github.com/spf13/cobra"

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(setCmd)
}
