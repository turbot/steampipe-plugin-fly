package fly

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/turbot/steampipe-plugin-fly/apiClient"
)

func getClient(ctx context.Context, d *plugin.QueryData) (*apiClient.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "fly"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*apiClient.Client), nil
	}

	// Prefer config options given in Steampipe
	flyConfig := GetConfig(d.Connection)
	if flyConfig.Token == nil {
		return nil, fmt.Errorf("token must be passed")
	}

	// Start with an empty Fly config
	config := apiClient.ClientConfig{Token: flyConfig.Token}

	// Create the client
	client, err := apiClient.CreateClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %s", err.Error())
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
