package ssh

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// RmCmd represents the {cred ssh rm <path>} command
var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files and directories",
	Long: `The rm command allows you to remove files and directories recursively.
It uses the 'rm' command to delete the specified path.

Examples:
  cred ssh rm <path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.RmLogic("ssh", args)
	},
}
