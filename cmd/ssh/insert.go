package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// InsertCmd represents the {cred ssh insert <filepath>} command
var InsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new ssh entry",
	Long: `The insert command allows you to add a new ssh entry to the ssh store.
You will be prompted to enter and confirm the ssh, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
ssh insert <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh insert <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>"

		sshStore := config.Constants.SshPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}

		// create a folder for the key
		path := args[0]
		keyFullPath := paths.BuildPath(sshStore, path)

		if paths.CheckPathExists(keyFullPath) {
			var choice string
			fmt.Print("The ssh key already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if choice != "y" && choice != "Y" {
				return
			}

			if err := os.RemoveAll(keyFullPath); err != nil {
				fmt.Println("failed to remove the file: ", err)
				return
			}
		}

		if err := os.MkdirAll(keyFullPath, 0700); err != nil {
			fmt.Println("failed to create ssh key folder: ", err)
			return
		}

		// get flags
		publicKeyPath, err := cmd.Flags().GetString("public-key")
		if err != nil {
			fmt.Println("failed to get public-key flag: ", err)
			return
		}

		privateKeyPath, err := cmd.Flags().GetString("private-key")
		if err != nil {
			fmt.Println("failed to get private-key flag: ", err)
			return
		}

		connectionString, err := cmd.Flags().GetString("connection-string")
		if err != nil {
			fmt.Println("failed to get connection-string flag: ", err)
			return
		}

		if publicKeyPath == "" && privateKeyPath == "" && connectionString == "" {
			fmt.Println("no files provided")
			return
		}

		// create files based on it
		files := make(map[string]string)
		files["public.gpg"] = publicKeyPath
		files["private.gpg"] = privateKeyPath

		for name, path := range files {
			if path == "" {
				continue
			}

			info, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("file %s does not exist\n", path)
				} else {
					fmt.Printf("failed to access file %s: %v\n", path, err)
				}
				removeCreated(keyFullPath)
				return
			}

			if info.IsDir() {
				fmt.Printf("path %s is a directory, not a file\n", path)
				removeCreated(keyFullPath)
				return
			}

			fullPath := paths.BuildPath(keyFullPath, name)
			content := ""
			file, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("failed to read file %s: %v\n", path, err)
				removeCreated(keyFullPath)
				return
			}
			content = string(file)

			if err := gpgcrypt.AddFile(fullPath, content, false); err != nil {
				fmt.Println("failed to insert ssh keys: ", err)
				removeCreated(keyFullPath)
				return
			}
		}

		if connectionString != "" {
			fullPath := paths.BuildPath(keyFullPath, "connection.gpg")
			connectionContent := connectionString

			if err := gpgcrypt.AddFile(fullPath, connectionContent, false); err != nil {
				fmt.Println("failed to insert ssh keys: ", err)
				removeCreated(keyFullPath)
				return
			}
		}

		fmt.Println("ssh inserted successfully")
	},
}

func init() {
	InsertCmd.Flags().String("public-key", "", "public key file path")
	InsertCmd.Flags().String("private-key", "", "private key file path")
	InsertCmd.Flags().String("connection-string", "", "connection string")
}

// function to remove created paths
func removeCreated(path string) {
	if err := os.RemoveAll(path); err != nil {
		fmt.Println("failed to remove the created files: ", err)
	}
}
