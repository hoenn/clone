package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	giturls "github.com/whilp/git-urls"
)

var rootCmd = &cobra.Command{
	Use:   "clone 'http(s) or git@ URL' [options]",
	Short: "Clone remote git repositories into a host/owner/repo file structure",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("expected exactly one URL argument")
			_ = cmd.Help()
			os.Exit(1)
		}
		rURL := args[0]
		u, err := giturls.Parse(rURL)
		if err != nil {
			fmt.Printf("Invalid URL '%s': %s\n", rURL, err)
			os.Exit(1)
		}
		fullPath := u.RequestURI()
		if strings.Contains(fullPath, "git@") {
			reqPath := strings.Split(u.RequestURI(), "/")
			fullPath = fmt.Sprintf("/%s/%s", reqPath[1], reqPath[2])

		}
		// make clone output relative
		fullPath = "." + fullPath
		fmt.Printf("Beginning clone of %s...\n", u.String())
		fmt.Println(fullPath)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
