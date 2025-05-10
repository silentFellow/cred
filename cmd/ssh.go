package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/cmd/ssh"
	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/git"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage SSH keys and connections",
	Long: `The ssh command allows you to manage SSH keys and establish connections.
It provides functionalities to add keys, list keys, and connect to servers.

Examples:
- Add a new SSH key: ssh add <key-name>
- List all SSH keys: ssh ls
- Connect to a server: ssh connect <key-name>`,

	Run: func(cmd *cobra.Command, args []string) {
		sshPath := config.Constants.SshPath

		if paths.CheckPathExists(sshPath) {
			err := utils.PrintTree(sshPath, "", true)
			if err != nil {
				fmt.Println("failed to parse ssh store: ", err)
			}
		}
	},
}

func init() {
	sshCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !gpgcrypt.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("GPG key not found, try [cred init <gpg-key-id>]")
		}

		if !gpgcrypt.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("invalid GPG key, try [cred init <gpg-key-id>]")
		}

		if !paths.CheckPathExists(config.Constants.SshPath) {
			err := os.MkdirAll(config.Constants.SshPath, 0700)
			if err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("failed to create ssh store: %v", err)
			}
		}

		return nil
	}

	sshCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	sshCmdsOnlyFiles := []*cobra.Command{
		ssh.ShowCmd,
	}

	for _, cmd := range sshCmdsOnlyFiles {
		cmd.ValidArgsFunction = fileCompletion(config.Constants.SshPath, false, true)
		sshCmd.AddCommand(cmd)
	}

	sshCmdsBoth := []*cobra.Command{
		ssh.InsertCmd,
		ssh.LsCmd,
		ssh.MkdirCmd,
		ssh.CpCmd,
	}

	for _, cmd := range sshCmdsBoth {
		cmd.ValidArgsFunction = fileCompletion(config.Constants.SshPath, true, true)
		sshCmd.AddCommand(cmd)
	}

	rootCmd.AddCommand(sshCmd)
}
