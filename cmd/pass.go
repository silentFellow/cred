package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/pass"
	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
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

		if utils.CheckPathExists(passPath) {
			err := utils.PrintTree(passPath, "", true)
			if err != nil {
				fmt.Printf("Failed to parse password store: %v\n", err)
			}
		}
	},
}

func init() {
	passCmd.AddCommand(pass.GenerateCmd)
	rootCmd.AddCommand(passCmd)
}
