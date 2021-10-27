package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clone 'URL' [options]",
	Short: "Clone remote git repositories into a host/owner/repo file structure",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("expected exactly one URL argument")
			_ = cmd.Help()
			os.Exit(1)
		}
		rURL := args[0]
		fmt.Printf("Beginning clone of %s...\n", rURL)
		fmt.Println(rURL)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
