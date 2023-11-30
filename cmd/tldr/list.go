package main

import (
	"fmt"
	"os"
	"strings"
)

func listPages(repoPath string, platform string) error {
	pagesDir := fmt.Sprintf(
		"%s/pages/%s",
		repoPath, platform,
	)

	entries, err := os.ReadDir(pagesDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(strings.TrimSuffix(entry.Name(), ".md"))
	}

	return nil
}
