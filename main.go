package main

import (
	"fmt"
	"os"

	"github.com/Pauloo27/tldr/repo"
	"github.com/Pauloo27/tldr/utils"
)

func main() {
	home, err := os.UserHomeDir()
	utils.HandleFatal("Cannot load user home", err)

	if len(os.Args) == 1 {
		fmt.Println("Missing page parameter")
		os.Exit(-1)
	} else {
		page := os.Args[1]
		err := repo.Load(home + "/.cache/tldr/")
		utils.HandleFatal("Cannot load repo", err)

		fmt.Println(page)
	}
}
