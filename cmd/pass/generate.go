package pass

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// GenerateCmd represents the {cred pass generate <filepath> [-l (length)]} command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new password and store it securely",
	Long: `The generate command creates a new password of specified length and stores it securely in the password store.
You can specify the length of the password using the -l flag. If the file already exists, you will be prompted to overwrite it.

Examples:
  pass generate mypassword -l 16
  pass generate anotherpassword -l 24`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass generate <filepath> [flags: -l {length}]"

		passStore := config.Constants.PassPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}
		filePath := args[0] + ".gpg"
		fullPath := paths.BuildPath(passStore, filePath)

		length, _ := cmd.Flags().GetInt("length")
		allowLower, _ := cmd.Flags().GetBool("allow-lower")
		allowUpper, _ := cmd.Flags().GetBool("allow-upper")
		allowDigit, _ := cmd.Flags().GetBool("allow-digit")
		allowSpecial, _ := cmd.Flags().GetBool("allow-special")
		generatedPassword := utils.GenerateRandom(
			length,
			allowLower,
			allowUpper,
			allowDigit,
			allowSpecial,
		)

		if paths.CheckPathExists(fullPath) {
			var choice string
			fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				return
			}

			if err := os.RemoveAll(fullPath); err != nil {
				fmt.Println("failed to remove the file: ", err)
			}
		}

		if err := gpgcrypt.AddFile(fullPath, generatedPassword, true); err != nil {
			fmt.Println("failed to insert password: ", err)
			return
		}
		fmt.Println("password inserted successfully and copied to clipboard")
	},
}

func init() {
	GenerateCmd.Flags().IntP("length", "l", 12, "length of the generated password")
	GenerateCmd.Flags().
		BoolP("allow-lower", "", true, "should allow lower-case characters in the password")
	GenerateCmd.Flags().
		BoolP("allow-upper", "", true, "should allow upper-case characters in the password")
	GenerateCmd.Flags().BoolP("allow-digit", "", true, "should allow digits in the password")
	GenerateCmd.Flags().
		BoolP("allow-special", "", true, "should allow special characters in the password")
}
