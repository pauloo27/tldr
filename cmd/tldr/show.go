package main

import (
	"fmt"
	"os"
)

func showPage(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}
