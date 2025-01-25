package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/internal/core"
)

// CopyCmd represents the {cred pass copy <filepath>} command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored password to system clipboard",
	Long: `The show command retrieves and copies the stored password for a given account,
Usage examples:

cred pass show <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CopyLogic("pass", args)
	},
}
