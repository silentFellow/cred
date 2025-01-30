package env

import (
	"github.com/silentFellow/cred/internal/core"
	"github.com/spf13/cobra"
)

// EditCmd represents the {cred env edit <filepath>} command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a env entry",
	Long: `Edit a env entry in the env store.

This command allows you to edit an existing env entry in your env store.
Usage:
  cred env edit <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.EditLogic("env", args)
	},
}
