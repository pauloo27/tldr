package repo

import (
	"os"
	"os/exec"
)

const (
	RemoteRepoURL = "https://github.com/tldr-pages/tldr.git"
)

func IsCloned(repoDir string) bool {
	_, err := os.Stat(repoDir)
	return !os.IsNotExist(err)
}

func Clone(repoDir string) error {
	cmd := exec.Command("git", "clone", RemoteRepoURL, repoDir)
	return cmd.Run()
}

func Pull(repoDir string) error {
	cmd := exec.Command("git", "-C", repoDir, "pull")
	return cmd.Run()
}
