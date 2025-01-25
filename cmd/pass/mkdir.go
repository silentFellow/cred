package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/internal/core"
)

// RmCmd represents the {cred pass rm <path>} command
var MkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Create directories",
	Long: `The mkdir command allows you to create directories, including nested directories.
It uses the 'mkdir' command to create the specified path.

Examples:
  cred pass mkdir <path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.MkdirLogic("pass", args)
	},
}
