package ssh

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// CopyCmd represents the {cred ssh copy filepath} command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored ssh to system clipboard",
	Long: `The show command retrieves and copies the stored ssh key for a given file.
Usage examples:

cred ssh copy <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CopyLogic("ssh", args)
	},
}
