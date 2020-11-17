package utils

import "log"

func HandleFatal(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
