package ssh

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/internal/core"
)

// ShowCmd represents the {cred ssh show <path>} command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the stored ssh",
	Long: `The show command retrieves and displays the stored ssh key for a given file.
Usage examples:

cred ssh show <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		core.ShowLogic("ssh", args)
	},
}
