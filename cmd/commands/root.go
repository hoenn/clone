package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
