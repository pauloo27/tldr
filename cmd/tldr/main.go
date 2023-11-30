package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	isUpdate  bool
	isVersion bool
	isList    bool
	language  string
	platform  string
	viewer    string
)

var rootCmd = &cobra.Command{
	Use:     "tldr [page]",
	Example: "tldr git clone",
	Short:   "TL;DR pages",
	Long:    "TL;DR pages reader written in Go, NOT FULLY complaint with the specification",
	Run:     handleCommand,
}

func init() {
	rootCmd.Flags().BoolVarP(
		&isVersion,
		"version", "v",
		false,
		"Shows the current version of the client, and the version of this specification that it implements",
	)

	rootCmd.Flags().BoolVarP(
		&isList,
		"list", "l",
		false,
		"List available pages for the current platform. Can be overriden with the --platform flag",
	)

	rootCmd.Flags().BoolVarP(
		&isUpdate,
		"update", "u",
		false,
		"Updates the offline cache of pages",
	)

	rootCmd.Flags().StringVarP(
		&platform,
		"platform", "p",
		"",
		"Specifies the platform to be used to perform the action (either listing or searching) as an argument. If not specified, the current one is used",
	)

	viper.SetConfigType("toml")
	viper.SetConfigName("config.toml")
	viper.AddConfigPath("$HOME/.config/tldr")
	viper.SetDefault("viewer", "less")
	viper.SetDefault("language", "en")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Cannot load config file: %s\n", err)
			os.Exit(1)
		}
	}

	defaultViewer := viper.GetString("viewer")
	defaultLanguage := viper.GetString("language")

	rootCmd.Flags().StringVarP(
		&language,
		"language", "L",
		defaultLanguage,
		"Specifies the preferred language for the page returned. Example: de or pt_BR. Against the spec, the LANG environment is ignored",
	)

	rootCmd.Flags().StringVarP(
		&viewer,
		"viewer", "V",
		defaultViewer,
		"Specifies viewer of the page. Example: less, bat, mdcat",
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
