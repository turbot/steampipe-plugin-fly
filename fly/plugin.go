package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-fly",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: errors.NotFoundError,
		},
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		TableMap: map[string]*plugin.Table{
			"fly_app":          tableFlyApp(ctx),
			"fly_machine":      tableFlyMachine(ctx),
			"fly_organization": tableFlyOrganization(ctx),
			"fly_volume":       tableFlyVolume(ctx),
		},
	}
	return p
}
