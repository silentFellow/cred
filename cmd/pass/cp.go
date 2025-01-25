package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/internal/core"
)

// CpCmd represents the {cred pass cp <src> <dest>} command
var CpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copies files and directories",
	Long: `The cp command allows you to copies files and directories.
It uses the 'cp' command to move the specified file or directory to a new location.

Examples:
  cred pass cp <source> <destination>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CpLogic("pass", args)
	},
}
