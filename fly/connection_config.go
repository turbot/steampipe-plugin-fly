package fly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type flyConfig struct {
	ApiToken *string `hcl:"api_token"`
}

func ConfigInstance() interface{} {
	return &flyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) flyConfig {
	if connection == nil || connection.Config == nil {
		return flyConfig{}
	}
	config, _ := connection.Config.(flyConfig)
	return config
}
