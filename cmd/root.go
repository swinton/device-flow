package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "device-flow",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(AuthCmd)
}
