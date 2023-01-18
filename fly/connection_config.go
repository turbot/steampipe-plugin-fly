package fly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type flyConfig struct {
	FlyApiToken *string `cty:"fly_api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"fly_api_token": {
		Type: schema.TypeString,
	},
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
