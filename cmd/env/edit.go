package env

import (
	"github.com/silentFellow/cred-store/cmd/common"
	"github.com/spf13/cobra"
)

// EditCmd represents the env/edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a env entry",
	Long: `Edit a env entry in the env store.

This command allows you to edit an existing env entry in your env store.
Usage:
  env edit <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.EditLogic("env", args)
	},
}
