package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Manage git repositories and operations",
	Long: `Manage git repositories and perform various git operations.
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
		gitCmd := exec.Command("git", "--version")
		if err := gitCmd.Run(); err != nil {
			cmd.SilenceUsage = true
			return fmt.Errorf("git is not installed")
		}

		return nil
	}

	rootCmd.AddCommand(gitCmd)
}
