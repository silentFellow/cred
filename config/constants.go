package config

import (
	"fmt"
	"os"
	"strings"
)

type constants struct {
	GpgKey    string
	Home      string
	Download  string
	Config    string
	StorePath string
	PassPath  string
	EnvPath   string
	Editor    string
	AutoGit   bool
}

var Constants constants = initConstants()

func initConstants() constants {
	home := getEnv("HOME", "/home")
	defaultConfig := fmt.Sprintf("%v/.config", home)

	gpgKey := ""
	if file, err := os.ReadFile(fmt.Sprintf("%v/.cred-store/.gpg-id", home)); err == nil {
		gpgKey = strings.TrimSpace(string(file))
	}

	return constants{
		GpgKey:    gpgKey,
		Home:      home,
		Config:    getEnv("XDG_CONFIG_HOME", defaultConfig),
		Download:  fmt.Sprintf("%v/Downloads", home),
		StorePath: fmt.Sprintf("%v/.cred-store", home),
		PassPath:  fmt.Sprintf("%v/.cred-store/pass", home),
		EnvPath:   fmt.Sprintf("%v/.cred-store/env", home),
		Editor:    getEnv("EDITOR", "vi"),
		AutoGit:   strings.ToLower(getEnv("CRED_STORE_AUTO_GIT", "false")) == "true",
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
