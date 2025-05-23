package ssh

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// MvCmd represents the {cred ssh mv <src> <dest>} command
var MvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move files and directories",
	Long: `The mv command allows you to move files and directories.
It uses the 'mv' command to move the specified file or directory to a new location.

Examples:
  cred ssh mv <source> <destination>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.MvLogic("ssh", args)
	},
}
