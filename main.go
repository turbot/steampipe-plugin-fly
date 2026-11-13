package main

import (
	"github.com/turbot/steampipe-plugin-fly/fly"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: fly.Plugin})
}
