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

	"github.com/silentFellow/cred/config"
)

func CopyToClipboard(text string, copyOnlyFirst bool) error {
	if copyOnlyFirst {
		text = strings.Split(text, "\n")[0]
	}

	if err := clipboard.WriteAll(text); err != nil {
		return err
	}

	return nil
}

// PrintTree recursively prints the directory structure in a tree-like format
func PrintTree(root string, prefix string, isLast bool) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}

	// if the path is a file
	// can't be open with tree
	if !info.IsDir() {
		fmt.Println(prefix + info.Name())
		return nil
	}

	// check if tree command present if so just execute it
	if config.Constants.Os != "windows" {
		treeCmd := SetCmd(CmdConfig{IsStdout: true}, "tree", root)
		if err := treeCmd.Run(); config.Constants.Os != "windows" && err == nil {
			return nil
		}
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
			return err
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

func GenerateRandom(
	n int,
	allowLower, allowUpper, allowDigit, allowSpecial bool,
	allowedSpecial string,
) string {
	getStringWithMinLen := func(s string, minLen int) string {
		var sBuilder strings.Builder
		sBuilder.Grow(minLen)
		sBuilder.WriteString(s)

		if sBuilder.Len() < minLen {
			sBuilder.WriteString(strings.Repeat(s, (minLen-sBuilder.Len())/len(s)+1))
		}

		return sBuilder.String()[:minLen]
	}

	var charsetBuilder strings.Builder

	lowerAllowed := "abcdefghijklmnopqrstuvwxyz"
	upperAllowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitAllowed := "0123456789"

	minLen := max(
		len(lowerAllowed),
		len(upperAllowed),
		len(digitAllowed),
		len(allowedSpecial),
	)

	if allowLower {
		charsetBuilder.WriteString(getStringWithMinLen(lowerAllowed, minLen))
	}
	if allowUpper {
		charsetBuilder.WriteString(getStringWithMinLen(upperAllowed, minLen))
	}
	if allowDigit {
		charsetBuilder.WriteString(getStringWithMinLen(digitAllowed, minLen))
	}
	if allowSpecial && len(allowedSpecial) > 0 {
		charsetBuilder.WriteString(getStringWithMinLen(allowedSpecial, minLen))
	}

	charset := charsetBuilder.String()

	var random strings.Builder
	for range n {
		random.WriteByte(charset[rand.IntN(len(charset))])
	}

	return random.String()
}

type CmdConfig struct {
	IsStdin  bool
	IsStdout bool
	IsStderr bool
	Dir      string
}

func SetCmd(cfg CmdConfig, args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	if strings.Trim(cfg.Dir, " ") != "" {
		cmd.Dir = cfg.Dir
	}

	if cfg.IsStdin {
		cmd.Stdin = os.Stdin
	}
	if cfg.IsStdout {
		cmd.Stdout = os.Stdout
	}
	if !config.Config.SuppressStderr && cfg.IsStderr {
		cmd.Stderr = os.Stderr
	}
	return cmd
}
