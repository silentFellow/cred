package env

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/common"
)

// ShowCmd represents the {cred env show} command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the stored env",
	Long: `The show command retrieves and displays the stored env for a given account.
Usage examples:

cred env show <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ShowLogic("env", args)
	},
}
