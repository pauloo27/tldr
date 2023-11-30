package main

import (
	"fmt"
	"os"
)

func findPage(repoPath, platform, language, page string) string {
	var pagesPath string

	if language == "en" {
		pagesPath = "pages"
	} else {
		pagesPath = fmt.Sprintf("pages.%s", language)
	}

	pagePath := fmt.Sprintf(
		"%s/%s/%s/%s.md",
		repoPath, pagesPath, platform, page,
	)

	if _, err := os.Stat(pagePath); os.IsNotExist(err) {
		if platform != "common" {
			return findPage(repoPath, "common", language, page)
		}
		return ""
	}

	return pagePath
}
