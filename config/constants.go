package config

import (
	"os"
	"runtime"
	"strings"

	"github.com/silentFellow/cred/internal/utils/paths"
)

type constants struct {
	GpgKey    string
	Home      string
	Download  string
	StorePath string
	PassPath  string
	EnvPath   string
	Os        string
}

var Constants constants = initConstants()

func initConstants() constants {
	// error won't occur, checks at cmd initiation
	// so program returns if err occurs
	home, _ := os.UserHomeDir()

	gpgKey := ""
	if file, err := os.ReadFile(paths.BuildPath(home, ".cred-store", ".gpg-id")); err == nil {
		gpgKey = strings.TrimSpace(string(file))
	}

	return constants{
		GpgKey:    gpgKey,
		Home:      home,
		Download:  paths.BuildPath(home, "Downloads"),
		StorePath: paths.BuildPath(home, ".cred-store"),
		PassPath:  paths.BuildPath(home, ".cred-store", "pass"),
		EnvPath:   paths.BuildPath(home, ".cred-store", "env"),
		Os:        runtime.GOOS,
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
