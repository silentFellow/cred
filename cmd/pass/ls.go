package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// LsCmd represents the {cred pass ls <path>} command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files and directories",
	Long: `The ls command allows you to list files and directories.
It uses the 'ls' command to display the contents of the current directory.

Examples:
  cred pass ls <path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.LsLogic("pass", args)
	},
}
