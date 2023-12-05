package main

import (
	"fmt"
	"os"
)

func printErrMsg(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func printErr(msg string, err error) {
	printErrMsg(fmt.Sprintf("%s: %s", msg, err))
}
