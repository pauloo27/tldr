package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	isUpdate  bool
	isVersion bool
	isList    bool   // TODO
	language  string // TODO
	platform  string // TODO
)

var rootCmd = &cobra.Command{
	Use:   "tldr [page]",
	Short: "TL;DR pages",
	Long:  "TL;DR pages reader written in Go, not fully complaint with the specification",
	Run:   handleCommand,
}

func init() {
	rootCmd.Flags().BoolVarP(
		&isVersion,
		"version", "v",
		false,
		"Shows the current version of the client, and the version of this specification that it implements",
	)

	rootCmd.Flags().BoolVarP(
		&isUpdate,
		"update", "u",
		false,
		"Updates the offline cache of pages",
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
