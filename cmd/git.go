package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
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
	Run: func(cmd *cobra.Command, args []string) {
		execCmd := exec.Command("git", args...)
		execCmd.Dir = config.Constants.StorePath
		output, _ := execCmd.CombinedOutput() // ignore the output else always status code throws
		fmt.Print(string(output))
	},
}

func init() {
	gitCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !git.CheckGitExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("git is not installed")
		}

		if !git.IsValidGitPath(config.Constants.StorePath) {
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

	rootCmd.AddCommand(gitCmd)
}
