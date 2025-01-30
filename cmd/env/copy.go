package env

import (
	"github.com/silentFellow/cred/internal/core"
	"github.com/spf13/cobra"
)

// CopyCmd represents the {env copy filepath} command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored env to system clipboard",
	Long: `The show command retrieves and copies the stored env for a given file.
Usage examples:

cred env copy <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CopyLogic("env", args)
	},
}
