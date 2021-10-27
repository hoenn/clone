package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	giturls "github.com/whilp/git-urls"
)

var (
	incHost, dryRun bool
)

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "show information but do not actually clone repository")
	rootCmd.Flags().BoolVarP(&incHost, "include-host", "i", true, "include host in clone structure")
}

var rootCmd = &cobra.Command{
	Use:   "clone 'http(s) or git@ URL'",
	Short: "Clone remote git repositories into a host/owner/repo file structure",
	Long:  "Clone remote git repositories into a host/owner/repo file structure relative to where this command is run",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
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
		if strings.Contains(rURL, "git@") {
			reqPath := strings.Split(u.RequestURI(), "/")
			fullPath = fmt.Sprintf("/%s/%s", reqPath[0], reqPath[1])
			fullPath = strings.Trim(fullPath, ".git")
		}
		fmt.Println(u.Host)
		if incHost {
			fullPath = fmt.Sprintf("/%s%s", u.Host, fullPath)
		}
		fmt.Println(fullPath)

		clonePath := "." + fullPath
		if dryRun {
			fmt.Println("Planning to clone:", rURL)
			fmt.Println("To:", clonePath)
			os.Exit(0)
		}

		fmt.Printf("Beginning clone of %s...\n", u.String())
		_, err = git.PlainClone(clonePath, false, &git.CloneOptions{
			URL:      rURL,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Printf("Failed to clone: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Repository cloned to: %s\n", clonePath)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
