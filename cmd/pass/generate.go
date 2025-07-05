package pass

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/core"
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
		allowLowerCase, _ := cmd.Flags().GetBool("allow-lowercase")
		allowUpperCase, _ := cmd.Flags().GetBool("allow-uppercase")
		allowDigit, _ := cmd.Flags().GetBool("allow-digit")
		allowSpecial, _ := cmd.Flags().GetBool("allow-special")
		allowedSpecial, _ := cmd.Flags().GetString("allowed-special")
		if !allowLowerCase && !allowUpperCase && !allowDigit && !allowSpecial {
			fmt.Println("failed to generate password: no character sets enabled")
			return
		}

		generatedPassword := utils.GenerateRandom(
			length,
			allowLowerCase,
			allowUpperCase,
			allowDigit,
			allowSpecial,
			allowedSpecial,
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
		fmt.Println("password generated successfully and copied to clipboard")

		isEditor, _ := cmd.Flags().GetBool("editor")
		if isEditor {
			core.EditLogic("pass", append([]string{filePath}, args[1:]...))
		}
	},
}

func init() {
	GenerateCmd.Flags().IntP("length", "l", 12, "length of the generated password")
	GenerateCmd.Flags().
		BoolP("allow-lowercase", "", true, "should allow lower-case characters in the password")
	GenerateCmd.Flags().
		BoolP("allow-uppercase", "", true, "should allow upper-case characters in the password")
	GenerateCmd.Flags().BoolP("allow-digit", "", true, "should allow digits in the password")
	GenerateCmd.Flags().
		BoolP("allow-special", "", true, "should allow special characters in the password")
	GenerateCmd.Flags().
		String("allowed-special", "!@#$%^&*()-_=+[]{}|;:,.<>?/`~", "allowed special characters in the password")
	GenerateCmd.Flags().
		BoolP("editor", "e", false, "open password in editor for editing extra details after insertion")
}
