package config

import (
	"fmt"
	"os"
)

const GH_HOST = "GH_HOST"

type Config struct {
	Host string
}

var config = Config{}

func init() {
	if host := os.Getenv(GH_HOST); host != "" {
		config.Host = host
	} else {
		config.Host = "github.com"
	}
}

func GetDeviceCodeURI() string {
	return fmt.Sprintf("https://%s/login/device/code", config.Host)
}

func GetAccessTokenURI() string {
	return fmt.Sprintf("https://%s/login/oauth/access_token", config.Host)
}
