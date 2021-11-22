package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	giturls "github.com/whilp/git-urls"
)

var (
	incHost, dryRun, progress bool
)

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "show information but do not actually clone repository.")
	rootCmd.Flags().BoolVar(&progress, "progress", true, "show progress of git to stdout.")
	rootCmd.Flags().BoolVarP(&incHost, "include-host", "i", false, "include host in cloned folder structure.")
}

var rootCmd = &cobra.Command{
	Use:   "clone 'http(s) or git@ URL'",
	Short: "Clone remote git repositories into a host/owner/repo file structure.",
	Long:  "Clone remote git repositories into a host/owner/repo file structure relative to where this command is run.",
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
		if incHost {
			fullPath = fmt.Sprintf("/%s%s", u.Host, fullPath)
		}

		clonePath := "." + fullPath
		if dryRun {
			fmt.Println("Planning to clone:", rURL)
			fmt.Println("To:", clonePath)
			os.Exit(0)
		}

		gitCmdArgs := []string{"clone"}
		if progress {
			gitCmdArgs = append(gitCmdArgs, "--progress")
		}

		gitCmdArgs = append(gitCmdArgs, rURL, clonePath)

		gitCmd := exec.Command("git", gitCmdArgs...)
		outPipe, err := gitCmd.StderrPipe()
		if err != nil {
			fmt.Printf("could not get std out pipe: %s\n", err)
			os.Exit(1)
		}
		err = gitCmd.Start()
		if err != nil {
			fmt.Printf("Failed to clone, status code from git: %s\n", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(outPipe)
		for scanner.Scan() {
			l := scanner.Text()
			fmt.Println(l)
		}
		gitCmd.Wait()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
