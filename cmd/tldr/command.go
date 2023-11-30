package main

import (
	"fmt"
	"os"

	"github.com/pauloo27/tldr/internal/log"
	"github.com/pauloo27/tldr/internal/repo"
	"github.com/spf13/cobra"
)

func handleCommand(cmd *cobra.Command, args []string) {
	if isVersion {
		printVersion()
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory", err)
		os.Exit(1)
	}

	repoPath := fmt.Sprintf("%s/.cache/tldr-repo", home)

	err = ensureRepoExists(repoPath)
	if err != nil {
		fmt.Println("Error cloning repository", err)
		os.Exit(1)
	}

	if isUpdate {
		fmt.Println("Updating repository...")
		err = repo.Pull(repoPath)
		if err != nil {
			fmt.Println("Error updating repository", err)
			os.Exit(1)
		}
		return
	}
}

func ensureRepoExists(repoPath string) error {
	if repo.IsCloned(repoPath) {
		return nil
	}
	log.PrintIfTTY(
		"Repository not found, cloning. It may take a while...",
	)

	return repo.Clone(repoPath)
}
