package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/env"
	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A command to manage env-variables",
	Long: `The pass command allows you to manage your env-variables efficiently.
It provides functionalities to create, update, and delete passwords.

Examples:
- Create a new env: env {insert/generate}
- Update an existing env: env edit
- Delete a env: env rm`,
	Run: func(cmd *cobra.Command, args []string) {
		envPath := config.Constants.EnvPath

		if utils.CheckPathExists(envPath) {
			err := utils.PrintTree(envPath, "", true)
			if err != nil {
				fmt.Printf("Failed to parse password store: %v\n", err)
			}
		}
	},
}

func init() {
	envCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !utils.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("GPG key not found, try [cred init <gpg-key-id>]")
		}

		if !utils.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("Invalid GPG key, try [cred init <gpg-key-id>]")
		}

		return nil
	}

	envCmd.AddCommand(env.InsertCmd)
	envCmd.AddCommand(env.CopyCmd)
	envCmd.AddCommand(env.ShowCmd)
	envCmd.AddCommand(env.EditCmd)
	envCmd.AddCommand(env.LsCmd)
	envCmd.AddCommand(env.RmCmd)
	envCmd.AddCommand(env.MkdirCmd)
	envCmd.AddCommand(env.MvCmd)
	rootCmd.AddCommand(envCmd)
}
