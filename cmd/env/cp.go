package env

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// CpCmd represents the {cred env cp <src> <dest>} command
var CpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copies files and directories",
	Long: `The cp command allows you to copies files and directories.
It uses the 'cp' command to move the specified file or directory to a new location.

Examples:
  cred env cp <source> <destination>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CpLogic("env", args)
	},
}
