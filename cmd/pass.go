package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/pass"
	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/git"
	"github.com/silentFellow/cred-store/internal/utils/paths"
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
				fmt.Printf("Failed to parse password store: %v\n", err)
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
			return fmt.Errorf("Invalid GPG key, try [cred init <gpg-key-id>]")
		}

		return nil
	}

	passCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	passCmd.AddCommand(pass.GenerateCmd)
	passCmd.AddCommand(pass.InsertCmd)
	passCmd.AddCommand(pass.ShowCmd)
	passCmd.AddCommand(pass.CopyCmd)
	passCmd.AddCommand(pass.EditCmd)
	passCmd.AddCommand(pass.LsCmd)
	passCmd.AddCommand(pass.RmCmd)
	passCmd.AddCommand(pass.MkdirCmd)
	passCmd.AddCommand(pass.MvCmd)
	passCmd.AddCommand(pass.CpCmd)
	rootCmd.AddCommand(passCmd)
}
