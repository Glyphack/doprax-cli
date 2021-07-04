package cmdutil

import (
	"fmt"

	"github.com/glyphack/doprax-cli/api"
	"github.com/spf13/viper"
)

type Factory struct {
	APIClient func() (*api.Client, error)
}

func New() *Factory {
	factory := Factory{APIClient: func() (*api.Client, error) {
		apiKey := viper.GetString("api-key")
		cca := viper.GetString("cca")

		if cca == "" || apiKey == "" {
			return nil, fmt.Errorf("Not logged in")
		}
		httpClient := api.NewHttpClient(&api.Config{ApiKey: apiKey, Cca: cca})
		return httpClient, nil
	}}

	return &factory
}
