package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-fly"

// Plugin creates this (fly) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: errors.NotFoundError,
		},
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		TableMap: map[string]*plugin.Table{
			"fly_app":                 tableFlyApp(ctx),
			"fly_app_certificate":     tableFlyAppCertificate(ctx),
			"fly_ip":                  tableFlyIP(ctx),
			"fly_location":            tableFlyLocation(ctx),
			"fly_machine":             tableFlyMachine(ctx),
			"fly_organization":        tableFlyOrganization(ctx),
			"fly_organization_member": tableFlyOrganizationMember(ctx),
			"fly_redis_database":      tableFlyRedisDatabase(ctx),
			"fly_volume":              tableFlyVolume(ctx),
		},
	}
	return p
}
