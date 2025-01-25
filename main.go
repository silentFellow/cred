package main

import (
	"fmt"
	"os/exec"

	"github.com/silentFellow/cred-store/cmd"
)

func main() {
	gpgCmd := exec.Command("gpg", "--version")
	if _, err := gpgCmd.CombinedOutput(); err != nil {
		fmt.Println("GPG is not installed. Please install GPG to use this tool")
		return
	}

	cmd.Execute()
}
