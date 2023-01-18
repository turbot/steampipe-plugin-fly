package fly

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
)

// getClient:: returns fly client after authentication
func getClient(ctx context.Context, d *plugin.QueryData) (*flyapi.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*flyapi.Client), nil
	}

	// Get the config
	flyConfig := GetConfig(d.Connection)

	/*
		precedence of credentials:
		- Credentials set in config
		- FLY_API_TOKEN env var
	*/
	var token string
	token = os.Getenv("FLY_API_TOKEN")

	if flyConfig.FlyApiToken != nil {
		token = *flyConfig.FlyApiToken
	}

	// Return if no credential specified
	if token == "" {
		return nil, fmt.Errorf("fly_api_token must be configured")
	}

	// Start with an empty Fly config
	config := flyapi.ClientConfig{FlyApiToken: flyConfig.FlyApiToken}

	// Create the client
	client, err := flyapi.CreateClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %s", err.Error())
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
