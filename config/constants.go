package config

import (
	"os"
	"strings"
)

type constants struct {
	GpgKey    string
	Home      string
	Download  string
	Config    config
	StorePath string
	PassPath  string
	EnvPath   string
	Editor    string
}

var Constants constants = initConstants()

func initConstants() constants {
	home := getEnv("HOME", "/home")

	configFilePath := buildPath(home, ".cred-store", "config")
	configMap := parseConfig(configFilePath)

	gpgKey := ""
	if file, err := os.ReadFile(buildPath(home, ".cred-store", ".gpg-id")); err == nil {
		gpgKey = strings.TrimSpace(string(file))
	}

	return constants{
		GpgKey:    gpgKey,
		Home:      home,
		Download:  buildPath(home, "Downloads"),
		StorePath: buildPath(home, ".cred-store"),
		PassPath:  buildPath(home, ".cred-store", "pass"),
		EnvPath:   buildPath(home, ".cred-store", "env"),
		Editor:    getEnv("EDITOR", "vi"),
		Config: config{
			Path: configFilePath,
			Values: configValues{
				AutoGit:        checkTrue(getConfigVal(configMap, "auto_git", "false")),
				SuppressStdout: checkTrue(getConfigVal(configMap, "suppress_stdout", "false")),
				SuppressStderr: checkTrue(getConfigVal(configMap, "suppress_stderr", "false")),
			},
		},
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func checkTrue(v string) bool {
	return strings.ToLower(strings.TrimSpace(v)) == "true"
}

// to avoid import cycles
func buildPath(pathSequence ...string) string {
	var path strings.Builder

	for i, currentPath := range pathSequence {
		path.WriteString(currentPath)

		if i != len(pathSequence)-1 && !strings.HasSuffix(currentPath, "/") {
			path.WriteString("/")
		}
	}

	return path.String()
}
