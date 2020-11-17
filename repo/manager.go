package repo

import (
	"fmt"
	"os"
	"os/exec"
)

var CATEGORIES = []string{"common", "linux", "osx", "sunos", "windows"}

const (
	REPO_PATH = "https://github.com/tldr-pages/tldr.git"
)

func Clone(repoDir string) (err error) {
	cmd := exec.Command("git", "clone", REPO_PATH, repoDir)
	err = cmd.Run()
	return
}

func Update(repoDir string) {
	// TODO
}

func IsCloned(repoDir string) bool {
	_, err := os.Stat(repoDir)
	return !os.IsNotExist(err)
}

func Load(repoDir string) error {
	if !IsCloned(repoDir) {
		fmt.Println("Repo not found, cloning to", repoDir)
		return Clone(repoDir)
	}
	return nil
}

func Find(repoPath, lang, pageName string) (path string) {
	langPath := "pages"
	if lang != "" {
		langPath += "." + lang
	}
	for _, category := range CATEGORIES {
		path = fmt.Sprintf("%s/%s/%s/%s.md", repoPath, langPath, category, pageName)
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			return
		}
	}
	path = ""
	return
}
