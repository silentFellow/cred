package utils

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(text string, copyOnlyFirst bool) error {
	if copyOnlyFirst {
		text = strings.Split(text, "\n")[0]
	}

	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("Failed to copy to clipboard: %w", err)
	}

	return nil
}

// PrintTree recursively prints the directory structure in a tree-like format
func PrintTree(root string, prefix string, isLast bool) error {
	info, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("error accessing %s: %w", root, err)
	}

	// check if tree command present if so just execute it
	treeCmd := SetCmd("", CmdIOConfig{IsStdout: true}, "tree", root)
	if err := treeCmd.Run(); err == nil {
		return nil
	}

	// Determine connector for the current item
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	fmt.Println(prefix + connector + info.Name())

	// If the current item is a directory, process its contents
	if info.IsDir() {
		entries, err := os.ReadDir(root)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %w", root, err)
		}

		// Sort entries alphabetically to maintain consistent output
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})

		// Iterate over directory entries
		for i, entry := range entries {
			isLastEntry := i == len(entries)-1
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			err := PrintTree(filepath.Join(root, entry.Name()), newPrefix, isLastEntry)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GenerateRandom(n int) string {
	var random strings.Builder
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"

	for range n {
		valType := rand.IntN(4)

		switch valType {
		case 0:
			random.WriteByte(byte('a' + rand.IntN(26)))
		case 1:
			random.WriteByte(byte('A' + rand.IntN(26)))
		case 2:
			random.WriteByte(byte('0' + rand.IntN(10)))
		case 3:
			random.WriteByte(specialChars[rand.IntN(len(specialChars))])
		}
	}

	return random.String()
}

type CmdIOConfig struct {
	IsStdin  bool
	IsStdout bool
	IsStderr bool
}

func SetCmd(filepath string, perm CmdIOConfig, args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	if strings.Trim(filepath, " ") != "" {
		cmd.Dir = filepath
	}

	if perm.IsStdin {
		cmd.Stdin = os.Stdin
	}
	if perm.IsStdout {
		cmd.Stdout = os.Stdout
	}
	if perm.IsStderr {
		cmd.Stderr = os.Stderr
	}
	return cmd
}
