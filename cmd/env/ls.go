package env

import (
	"github.com/silentFellow/cred/internal/core"
	"github.com/spf13/cobra"
)

// LsCmd represents the {cred env ls <path>} command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files and directories",
	Long: `The ls command allows you to list files and directories.
It uses the 'ls' command to display the contents of the current directory.

Examples:
  cred env ls <path>`,
	Run: func(cmd *cobra.Command, args []string) {
    core.LsLogic("env", args)
	},
}
