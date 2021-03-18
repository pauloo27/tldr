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

func Update(repoDir string) (err error) {
	cmd := exec.Command("git", "-C", repoDir, "pull")
	err = cmd.Run()
	return
}

func IsCloned(repoDir string) bool {
	_, err := os.Stat(repoDir)
	return !os.IsNotExist(err)
}

func Load(repoDir string) error {
	if !IsCloned(repoDir) {
		fmt.Printf("Repo not found, cloning to %s... It may take a while...\n", repoDir)
		return Clone(repoDir)
	}
	return nil
}

func Find(repoPath, lang, pageName string) (string, error) {
	path := ""
	langPath := "pages"
	if lang != "" {
		langPath += "." + lang
	}
	for _, category := range CATEGORIES {
		path = fmt.Sprintf("%s/%s/%s/%s.md", repoPath, langPath, category, pageName)
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			return path, nil
		}
	}
	err := fmt.Errorf(
		"Cannot find page %s in the lang %s in any of the categories %v",
		pageName,
		lang,
		CATEGORIES,
	)
	return "", err
}
