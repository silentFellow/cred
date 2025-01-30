package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/cmd/pass"
	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/completions"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/git"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// passCmd represents the pass command
var passCmd = &cobra.Command{
	Use:   "pass",
	Short: "A command to manage passwords",
	Long: `The pass command allows you to manage your passwords efficiently.
It provides functionalities to create, update, and delete passwords.

Examples:
- Create a new password: pass {insert/generate}
- Update an existing password: pass edit
- Delete a password: pass rm`,
	Run: func(cmd *cobra.Command, args []string) {
		passPath := config.Constants.PassPath

		if paths.CheckPathExists(passPath) {
			err := utils.PrintTree(passPath, "", true)
			if err != nil {
				fmt.Println("failed to parse password store: ", err)
			}
		}
	},
}

func init() {
	passCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !gpgcrypt.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("GPG key not found, try [cred init <gpg-key-id>]")
		}

		if !gpgcrypt.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("invalid GPG key, try [cred init <gpg-key-id>]")
		}

		if !paths.CheckPathExists(config.Constants.PassPath) {
			err := os.MkdirAll(config.Constants.PassPath, 0700)
			if err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("failed to create pass store: %v", err)
			}
		}

		return nil
	}

	passCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	passCmds := []*cobra.Command{
		pass.GenerateCmd,
		pass.InsertCmd,
		pass.ShowCmd,
		pass.CopyCmd,
		pass.EditCmd,
		pass.MkdirCmd,
		pass.LsCmd,
		pass.RmCmd,
		pass.MvCmd,
		pass.CpCmd,
	}

	for _, cmd := range passCmds {
		cmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			suggestions := completions.GetFilePathSuggestions(
				config.Constants.PassPath,
			)

			// If no suggestions are found, return an empty slice
			if len(suggestions) == 0 {
				return []string{}, cobra.ShellCompDirectiveNoFileComp
			}

			return suggestions, cobra.ShellCompDirectiveDefault
		}

		passCmd.AddCommand(cmd)
	}

	rootCmd.AddCommand(passCmd)
}
