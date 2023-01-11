package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFlyApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_app",
		Description: "Fly App",
		List: &plugin.ListConfig{
			Hydrate: listFlyApps,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyApp,
			KeyColumns: plugin.SingleColumn("name"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "network",
				Description: "Specifies the app network.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "app_url",
				Description: "The URL of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hostname",
				Description: "The hostname of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "deployed",
				Description: "If true, the app is successfully deployed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "current_release_id",
				Description: "Specifies the ID of the current app release.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CurrentRelease.Id"),
			},
			{
				Name:        "organization",
				Description: "Specifies the organization details where the app is deployed.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "autoscaling",
				Description: "Specifies the autoscaling information.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "config",
				Description: "Specifies the app configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "health_checks",
				Description: "Specifies the app health check information.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ip_addresses",
				Description: "A list of IP addresses associated with the app.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listFlyApps(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app.listFlyApps", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListAppsRequestConfiguration{}

	// There is no max page limit as such defined
	// but, we set the default page limit as 5000
	pageLimit := 5000

	// Adjust page limit, if less than default value
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < pageLimit {
			pageLimit = int(*limit)
		}
	}
	options.Limit = pageLimit

	for {
		query, err := flyapi.ListApps(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_app.listFlyApps", "query_error", err)
			return nil, err
		}

		for _, app := range query.Apps.Nodes {
			d.StreamListItem(ctx, app)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.Apps.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.Apps.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")
	if name == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app.getFlyApp", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetFullApp(context.Background(), conn.Graphql, name)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app.getFlyApp", "query_error", err)
		return nil, err
	}

	return query.App, nil
}
