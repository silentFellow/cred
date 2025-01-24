/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

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
	rootCmd.AddCommand(envCmd)
}
