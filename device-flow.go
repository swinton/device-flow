package main

import (
	"fmt"
	"net/http"

	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
)

func main() {
	var scopes []string

	var rootCmd = &cobra.Command{
		Use:       "[client_id] [-S scope]...",
		Short:     "Generates an access token using the device flow",
		ValidArgs: []string{"client_id"},
		Args:      cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			clientID := args[0]
			httpClient := http.DefaultClient

			code, err := device.RequestCode(httpClient, "https://github.com/login/device/code", clientID, scopes)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Copy code: %s\n", code.UserCode)
			fmt.Printf("then open: %s\n", code.VerificationURI)

			accessToken, err := device.PollToken(httpClient, "https://github.com/login/oauth/access_token", clientID, code)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Access token: %s\n", accessToken.Token)
		},
	}
	rootCmd.Flags().StringArrayVarP(&scopes, "scopes", "S", []string{}, "scopes to request through device flow, e.g. repo, read:org")
	rootCmd.Execute()
}
