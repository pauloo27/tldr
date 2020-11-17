package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/Pauloo27/tldr/repo"
	"github.com/Pauloo27/tldr/utils"
	"github.com/joho/godotenv"
)

func main() {
	home, err := os.UserHomeDir()
	utils.HandleFatal("Cannot load user home", err)

	repoPath := home + "/.cache/tldr"

	err = godotenv.Load()
	utils.HandleFatal("Cannot load .env", err)

	lang := os.Getenv("TLDR_LANG")
	viewer := os.Getenv("TLDR_VIEWER")

	if len(os.Args) == 1 {
		fmt.Println("Missing page parameter")
		os.Exit(-1)
	} else {
		pageName := os.Args[1]
		err := repo.Load(repoPath)
		utils.HandleFatal("Cannot load repo", err)

		pagePath, err := repo.Find(repoPath, lang, pageName)

		utils.HandleFatal("Cannot find page", err)

		err = syscall.Exec(viewer, []string{viewer, pagePath}, os.Environ())
		utils.HandleFatal("Cannot run viewer", err)
	}
}
