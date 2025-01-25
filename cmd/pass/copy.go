package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/common"
)

// CopyCmd represents the pass/show command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored password to system clipboard",
	Long: `The show command retrieves and copies the stored password for a given account.
Usage examples:

pass show <account_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.CopyLogic("pass", args)
	},
}
