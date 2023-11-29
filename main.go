package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/pauloo27/tldr/repo"
	"github.com/pauloo27/tldr/utils"
)

func main() {
	home, err := os.UserHomeDir()
	utils.HandleFatal("Cannot load user home", err)

	repoPath := home + "/.cache/tldr"

	dotenvPath := home + "/.config/tldr.env"

	err = godotenv.Load(dotenvPath)
	utils.HandleFatal("Cannot load .env. Copy the default .env to "+dotenvPath, err)

	lang := os.Getenv("TLDR_LANG")
	viewer := os.Getenv("TLDR_VIEWER")

	if len(os.Args) == 1 {
		fmt.Println("Missing page parameter")
		os.Exit(-1)
	} else {
		err := repo.Load(repoPath)
		utils.HandleFatal("Cannot load repo", err)

		pageName := os.Args[1]
		if pageName == "-u" {
			err = repo.Update(repoPath)
			utils.HandleFatal("Cannot update repo", err)
			fmt.Println("Cache updated.")
			return
		}

		pagePath, err := repo.Find(repoPath, lang, pageName)
		utils.HandleFatal("Cannot find page", err)

		err = syscall.Exec(viewer, []string{viewer, pagePath}, os.Environ())
		utils.HandleFatal("Cannot run viewer", err)
	}
}
