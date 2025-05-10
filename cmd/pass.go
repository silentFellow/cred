package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/cmd/pass"
	"github.com/silentFellow/cred/config"
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

	passCmdsOnlyFiles := []*cobra.Command{
		pass.ShowCmd,
		pass.CopyCmd,
		pass.EditCmd,
	}

	for _, cmd := range passCmdsOnlyFiles {
		cmd.ValidArgsFunction = fileCompletion(config.Constants.PassPath, false, true)
		passCmd.AddCommand(cmd)
	}

	passCmdsBoth := []*cobra.Command{
		pass.GenerateCmd,
		pass.InsertCmd,
		pass.MkdirCmd,
		pass.LsCmd,
		pass.RmCmd,
		pass.MvCmd,
		pass.CpCmd,
	}

	for _, cmd := range passCmdsBoth {
		cmd.ValidArgsFunction = fileCompletion(config.Constants.PassPath, true, true)
		passCmd.AddCommand(cmd)
	}

	rootCmd.AddCommand(passCmd)
}
