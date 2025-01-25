package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/common"
)

// EditCmd represents the pass/edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a password entry",
	Long: `Edit a password entry in the password store.

This command allows you to edit an existing password entry in your password store.
Usage:
  pass edit <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.EditLogic("pass", args)
	},
}
