package pass

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/internal/core"
)

// ShowCmd represents the {cred pass show <account_path>} command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the stored password",
	Long: `The show command retrieves and displays the stored password for a given account.
Usage examples:

cred pass show <account_path>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.ShowLogic("pass", args)
	},
}
