package main

import "fmt"

const (
	appVersion  = "0.0.1"
	specVersion = "2.0"
)

func printVersion() {
	fmt.Printf("Client version: %s\n", appVersion)
	fmt.Printf("Spec version: %s\n", specVersion)
}
