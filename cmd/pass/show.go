package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/common"
)

// pass/showCmd represents the pass/show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the stored password",
	Long: `The show command retrieves and displays the stored password for a given account.
Usage examples:

pass show <account_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ShowLogic("pass", args)
	},
}
