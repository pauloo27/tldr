package xdg

import (
	"os"
	"path/filepath"
)

func ConfigDir() string {
	if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
		return filepath.Join(xdgConfig, "tldr")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "tldr")
}

func CacheDir() string {
	if xdgCache := os.Getenv("XDG_CACHE_HOME"); xdgCache != "" {
		return filepath.Join(xdgCache, "tldr")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cache", "tldr")
}
