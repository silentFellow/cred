package main

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred-store/cmd"
	"github.com/silentFellow/cred-store/internal/utils"
)

func main() {
	gpgCmd := utils.SetCmd("", utils.CmdIOConfig{}, "gpg", "--version")
	if _, err := gpgCmd.CombinedOutput(); err != nil {
		fmt.Println("GPG is not installed. Please install GPG to use this tool")
		return
	}

	if _, err := os.UserHomeDir(); err != nil {
		fmt.Println(
			"Home directory not found. Please ensure your home directory is set up correctly",
		)
		return
	}

	cmd.Execute()
}
