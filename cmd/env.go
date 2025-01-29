package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/env"
	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/completions"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/git"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A command to manage env-variables",
	Long: `The env command allows you to manage your env-variables efficiently.
It provides functionalities to create, update, and delete env-variables.

Examples:
- Create a new env: env {insert/generate}
- Update an existing env: env edit
- Delete a env: env rm`,
	Run: func(cmd *cobra.Command, args []string) {
		envPath := config.Constants.EnvPath

		if paths.CheckPathExists(envPath) {
			err := utils.PrintTree(envPath, "", true)
			if err != nil {
				fmt.Printf("failed to parse env store: %v\n", err)
			}
		}
	},
}

func init() {
	envCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !gpgcrypt.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("gpg key not found, try [cred init <gpg-key-id>]")
		}

		if !gpgcrypt.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("invalid gpg key, try [cred init <gpg-key-id>]")
		}

		if !paths.CheckPathExists(config.Constants.EnvPath) {
			err := os.MkdirAll(config.Constants.EnvPath, 0700)
			if err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("failed to create env store: %v", err)
			}
		}

		return nil
	}

	envCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	envCmds := []*cobra.Command{
		env.InsertCmd,
		env.CopyCmd,
		env.ShowCmd,
		env.EditCmd,
		env.LsCmd,
		env.RmCmd,
		env.MkdirCmd,
		env.MvCmd,
		env.CpCmd,
	}

	for _, cmd := range envCmds {
		cmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			suggestions := completions.GetFilePathSuggestions(
				config.Constants.EnvPath,
			)

			// If no suggestions are found, return an empty slice
			if len(suggestions) == 0 {
				return []string{}, cobra.ShellCompDirectiveNoFileComp
			}

			return suggestions, cobra.ShellCompDirectiveDefault
		}

		envCmd.AddCommand(cmd)
	}

	envCmd.AddCommand(env.GetCmd)
	rootCmd.AddCommand(envCmd)
}
