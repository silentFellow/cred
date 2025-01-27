package config

import (
	"os"
	"strings"
)

type configValues struct {
	AutoGit        bool
	SuppressStdout bool
	SuppressStderr bool
}

type config struct {
	Path   string
	Values configValues
}

type configMapType map[string]string

func parseConfig(filePath string) configMapType {
	configMap := make(map[string]string)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return configMap
	}

	output := string(file)
	formattedOutput := strings.ReplaceAll(output, " ", "")

	for _, line := range strings.Split(formattedOutput, "\n") {
		entry := strings.Split(line, "=")
		if len(entry) == 2 {
			configMap[entry[0]] = entry[1]
		}
	}

	return configMap
}

func getConfigVal(configMap configMapType, key, fallback string) string {
	if val, ok := configMap[key]; ok {
		return val
	}

	return fallback
}
