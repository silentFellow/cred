package config

import (
	"fmt"
	"os"
)

type constants struct {
  GpgKey string
  Home string
  Config string
  StorePath string
  PassPath string
  EnvPath string
  Editor string
}

var Constants constants = initConstants()

func initConstants() constants {
  home := getEnv("HOME", "/home")
  defaultConfig := fmt.Sprintf("%v/.config", home)

  gpgKey := ""
  if file, err := os.ReadFile(fmt.Sprintf("%v/.cred-store/.gpg-id", home)); err == nil {
    gpgKey = string(file)
  }

  return constants {
    GpgKey: gpgKey,
    Home: home,
    Config: getEnv("XDG_CONFIG_HOME", defaultConfig),
    StorePath: fmt.Sprintf("%v/.cred-store", home),
    PassPath: fmt.Sprintf("%v/.cred-store/pass", home),
    EnvPath: fmt.Sprintf("%v/.cred-store/env", home),
    Editor: getEnv("EDITOR", "vi"),
  }
}

func getEnv(key, fallback string) string {
  if val, ok := os.LookupEnv(key); ok {
    return val
  }

  return fallback
}
