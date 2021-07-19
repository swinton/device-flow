package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "device-flow",
	}
)

func Factory(version string) *cobra.Command {
	rootCmd.Version = version
	return rootCmd
}

func init() {
	rootCmd.AddCommand(AuthCmd)
}
