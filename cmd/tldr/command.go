package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pauloo27/tldr/internal/repo"
	"github.com/pauloo27/tldr/internal/tty"
	"github.com/spf13/cobra"
)

func handleCommand(cmd *cobra.Command, args []string) int {
	if isVersion {
		printVersion()
		return 0
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory", err)
		return 1
	}

	repoPath := fmt.Sprintf("%s/.cache/tldr-repo", home)

	err = ensureRepoExists(repoPath)
	if err != nil {
		fmt.Println("Error cloning repository", err)
		return 1
	}

	if isUpdate {
		fmt.Println("Updating repository...")
		err = repo.Pull(repoPath)
		if err != nil {
			fmt.Println("Error updating repository", err)
			return 1
		}
		return 0
	}

	if platform == "" {
		osName := runtime.GOOS
		switch osName {
		case "darwin":
			platform = "osx"
		case "linux", "windows", "freebsd", "openbsd", "netbsd", "sunos":
			platform = osName
		default:
			fmt.Printf("Platform %s not supported\n", osName)
			return 1
		}
	}

	if isList {
		err := listPages(repoPath, platform)
		if err != nil {
			fmt.Println("Error listing pages", err)
			return 1
		}
		return 0
	}

	if len(args) == 0 {
		fmt.Println("No page specified")
		_ = cmd.Help()
		return 1
	}

	pageStrBuilder := strings.Builder{}

	for i, arg := range args {
		if i != 0 {
			pageStrBuilder.WriteString("-")
		}
		pageStrBuilder.WriteString(strings.ToLower(arg))
	}

	page := pageStrBuilder.String()

	path := findPage(repoPath, platform, language, page)

	if path == "" {
		fmt.Println("Page not found. Maybe try updating the cache with --update?")
		return 1
	}

	err = showPage(path)
	if err != nil {
		fmt.Println("Error showing page", err)
		return 1
	}

	return 0
}

func ensureRepoExists(repoPath string) error {
	if repo.IsCloned(repoPath) {
		return nil
	}
	if tty.IsTTY() {
		fmt.Println(
			"Repository not found, cloning. It may take a while...",
		)
	}

	return repo.Clone(repoPath)
}
