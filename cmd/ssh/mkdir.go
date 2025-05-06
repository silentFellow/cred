package ssh

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// RmCmd represents the {cred ssh rm <path>} command
var MkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Create directories",
	Long: `The mkdir command allows you to create directories, including nested directories.
It uses the 'mkdir' command to create the specified path.

Examples:
  cred ssh mkdir <path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.MkdirLogic("ssh", args)
	},
}
