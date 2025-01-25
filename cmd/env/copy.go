package env

import (
	"github.com/silentFellow/cred-store/cmd/common"
	"github.com/spf13/cobra"
)

// CopyCmd represents the env/copy command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored env to system clipboard",
	Long: `The show command retrieves and copies the stored env for a given account.
Usage examples:

cred copy show <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.CopyLogic("env", args)
	},
}
