package pass

import (
	"github.com/silentFellow/cred/internal/core"
	"github.com/spf13/cobra"
)

// EditCmd represents the {cred pass edit <filepath>} command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a password entry",
	Long: `Edit a password entry in the password store.

This command allows you to edit an existing password entry in your password store.
Usage:
  cred pass edit <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.EditLogic("pass", args)
	},
}
