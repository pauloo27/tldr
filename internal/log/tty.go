package log

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func PrintIfTTY(line string) {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println(line)
	}
}
