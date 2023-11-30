package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pauloo27/tldr/internal/tty"
)

func showPage(path string) error {
	if !tty.IsTTY() {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	}

	cmd := exec.Command(viewer, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
