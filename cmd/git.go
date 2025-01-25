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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
