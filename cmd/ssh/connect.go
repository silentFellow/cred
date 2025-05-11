package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
	sshUtils "github.com/silentFellow/cred/internal/utils/ssh"
)

// ConnectCmd represents the {cred ssh connect <filepath>} command
var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Establish SSH connection using stored connection string and private key",
	Long: `Looks for the 'connection.gpg' and 'private.gpg' files stored under the given <keyname>
in the cred store. Decrypts these files and uses the connection string and private key
to establish an SSH connection to the remote server.

Example:
  cred ssh connect my-server
This will look for:
  my-server/connection.gpg
  my-server/private.gpg
and use the decrypted contents to connect.`,

	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh connect <key-name>"

		sshStore := config.Constants.SshPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}

		path := args[0]
		keyFullPath := paths.BuildPath(sshStore, path)

		if !paths.CheckPathExists(keyFullPath) {
			fmt.Println("SSH key entry does not exist")
			return
		}

		keyInfo, err := os.Stat(keyFullPath)
		if err != nil {
			fmt.Println("failed to retrieve SSH key info: ", err)
			return
		}

		if !keyInfo.IsDir() {
			fmt.Println("specified SSH key path must be a directory")
			return
		}

		connectionFile := paths.BuildPath(keyFullPath, "connection.gpg")
		connectionFileInfo, err := os.Stat(connectionFile)
		if err != nil {
			fmt.Println("failed to retrieve connection file info: ", err)
			return
		}

		if connectionFileInfo.IsDir() {
			fmt.Println("specified connection file path must be a file")
		}

		privateFile := paths.BuildPath(keyFullPath, "private.gpg")
		privateFileInfo, err := os.Stat(privateFile)
		if err != nil {
			fmt.Println("failed to retrieve private file info: ", err)
			return
		}

		if privateFileInfo.IsDir() {
			fmt.Println("specified private file path must be a file")
			return
		}

		connectionContent, err := gpgcrypt.Decrypt(connectionFile)
		if err != nil {
			fmt.Println("failed to decrypt connection file: ", err)
			return
		}

		privateKeyFileContent, err := gpgcrypt.Decrypt(privateFile)
		if err != nil {
			fmt.Println("failed to decrypt private file: ", err)
			return
		}

		// create a temp file to store the gpg decrypted private file value
		sshFile, err := os.CreateTemp("", "ssh_private_key_*")
		if err != nil {
			fmt.Println("failed to create a temp file to store decrypted ssh key: ", err)
			return
		}
		defer func() {
			if err := os.Remove(sshFile.Name()); err != nil {
				fmt.Printf(
					"failed to remove the ssh file after used(remove manually at %v): %v\n",
					sshFile.Name(),
					err,
				)
			}
		}()
		if _, err := sshFile.WriteString(privateKeyFileContent); err != nil {
			fmt.Println("failed to write the decrypted SSH key to the temporary file: ", err)
			return
		}

		connectionCmd := utils.SetCmd(
			utils.CmdConfig{
				IsStdin:  true,
				IsStdout: true,
				IsStderr: true,
			},
			"ssh",
			"-i",
			sshFile.Name(),
			connectionContent,
		)

		if err := connectionCmd.Run(); err != nil {
			fmt.Println("failed to run ssh command: ", err)
			return
		}
	},
}

func init() {
	ConnectCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !sshUtils.CheckSshExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("SSH is not installed")
		}

		return nil
	}
}
