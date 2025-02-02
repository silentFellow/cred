package config

import (
	"os"
	"strings"

	"github.com/silentFellow/cred/internal/utils/paths"
)

type config struct {
	AutoGit        bool
	SuppressStderr bool
	Editor         string
}

var (
	Config     config = initConfig()
	ConfigPath string = paths.BuildPath(Constants.Home, ".cred-store", "config")
)

func initConfig() config {
	configMap := parseConfig(ConfigPath)

	return config{
		AutoGit:        checkTrue(getConfigVal(configMap, "auto_git", "false")),
		SuppressStderr: checkTrue(getConfigVal(configMap, "suppress_stderr", "false")),
		Editor:         getConfigVal(configMap, "editor", getDefaultEditor()),
	}
}

func parseConfig(filePath string) map[string]string {
	configMap := make(map[string]string)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return configMap
	}

	output := string(file)
	formattedOutput := strings.ReplaceAll(output, " ", "")

	for _, line := range strings.Split(formattedOutput, "\n") {
		if strings.HasPrefix(line, "#") { // comments
			continue
		}

		entry := strings.Split(line, "=")
		if len(entry) == 2 {
			configMap[entry[0]] = entry[1]
		}
	}

	return configMap
}

func getConfigVal(configMap map[string]string, key, fallback string) string {
	if val, ok := configMap[key]; ok {
		return val
	}

	return fallback
}

func checkTrue(v string) bool {
	return strings.ToLower(strings.TrimSpace(v)) == "true"
}

func getDefaultEditor() string {
	if Constants.Os == "windows" {
		return "notepad"
	}

	return "vi"
}
