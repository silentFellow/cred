package config

import (
	"fmt"
	"os"
)

type constants struct {
  Home string
  Config string
  StorePath string
  PassPath string
  EnvPath string
}

var Constants constants = initConstants()

func initConstants() constants {
  home := getEnv("HOME", "/home")
  defaultConfig := fmt.Sprintf("%v/.config", home)

  return constants {
    Home: home,
    Config: getEnv("XDG_CONFIG_HOME", defaultConfig),
    StorePath: fmt.Sprintf("%v/.cred-store", home),
    PassPath: fmt.Sprintf("%v/.cred-store/pass", home),
    EnvPath: fmt.Sprintf("%v/.cred-store/env", home),
  }
}

func getEnv(key, fallback string) string {
  if val, ok := os.LookupEnv(key); ok {
    return val
  }

  return fallback
}
