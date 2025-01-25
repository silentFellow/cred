package env

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/internal/core"
)

// RmCmd represents the {cred env rm <path>} command
var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files and directories",
	Long: `The rm command allows you to remove files and directories recursively.
It uses the 'rm' command to delete the specified path.

Examples:
  cred env rm <path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.RmLogic("env", args)
	},
}
