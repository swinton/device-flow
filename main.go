package main

import (
	"fmt"
	"os"

	"github.com/swinton/device-flow/cmd"
)

var version = "dev"

func main() {
	rootCmd := cmd.Factory(version)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
