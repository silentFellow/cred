package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/completions"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/git"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Manage cred-store git repository and operations",
	Long: `Manage cred-store git repository and perform various git operations.
This command allows you to interact with git repositories, perform updates,
and manage your version control workflow. For example:

cred git <command> [arguments]`,
}

func init() {
	gitCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !git.CheckGitExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("git is not installed")
		}

		if !(strings.ToLower(cmd.Name()) == "init") && !git.IsValidGitPath(config.Constants.StorePath) {
			var choice string
			fmt.Print(
				"github repository not found in the store path. Do you want to initialize a new repository? (y/n): ",
			)
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				cmd.SilenceUsage = true
				return fmt.Errorf("git repository not found")
			}

			if err := git.InitRepo(config.Constants.StorePath); err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("failed to initialize git repository: %w", err)
			}
		}

		return nil
	}

	// Add subcommands for each Git command
	for cmd, desc := range completions.GitCommandMap {
		subCmd := &cobra.Command{
			Use:                cmd,
			Short:              desc,
			DisableFlagParsing: true, // to avoid parsing flags for git commands
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				cmdString := append([]string{"git", cmd.Use}, args...)
				execCmd := utils.SetCmd(
					config.Constants.StorePath,
					utils.CmdIOConfig{IsStdin: true, IsStdout: true, IsStderr: true},
					cmdString...,
				)

				if err := execCmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "Error running git command: %v\n", err)
					os.Exit(1)
				}
			},
			ValidArgsFunction: completions.GetGitFileCompletion,
		}

		gitCmd.AddCommand(subCmd)
	}

	rootCmd.AddCommand(gitCmd)
}
