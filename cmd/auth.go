package cmd

import (
	"fmt"
	"net/http"

	"github.com/cli/oauth/device"
	"github.com/spf13/cobra"
	"github.com/swinton/device-flow/cmd/internal/config"
)

var (
	AuthCmd = &cobra.Command{
		Use:       "auth [client_id] [-S scope]...",
		Short:     "Generate an access token using the OAuth device authorization flow",
		ValidArgs: []string{"client_id"},
		Args:      cobra.ExactArgs(1),
		RunE:      run,
	}
	scopes []string
)

func init() {
	AuthCmd.Flags().StringArrayVarP(&scopes, "scopes", "S", []string{}, "scopes to request through device flow, e.g. repo, read:org")
}

func run(cmd *cobra.Command, args []string) error {
	clientID := args[0]
	httpClient := http.DefaultClient

	code, err := device.RequestCode(httpClient, config.GetDeviceCodeURI(), clientID, scopes)
	if err != nil {
		return err
	}

	fmt.Printf("Copy code: %s\n", code.UserCode)
	fmt.Printf("then open: %s\n", code.VerificationURI)

	accessToken, err := device.PollToken(httpClient, config.GetAccessTokenURI(), clientID, code)
	if err != nil {
		return err
	}

	fmt.Printf("Access token: %s\n", accessToken.Token)
	fmt.Printf("Refresh token: %s\n", accessToken.RefreshToken)

	return nil
}
