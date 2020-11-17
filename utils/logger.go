package utils

import (
	"fmt"
	"os"
)

func HandleFatal(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
		os.Exit(-1)
	}
}
