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
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the app."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the app."},
			{Name: "network", Type: proto.ColumnType_STRING, Description: "Specifies the app network."},
			{Name: "app_url", Type: proto.ColumnType_STRING, Description: "The URL of the app."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The hostname of the app."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the app."},
			{Name: "deployed", Type: proto.ColumnType_BOOL, Description: "If true, the app is successfully deployed."},
			{Name: "current_release_id", Type: proto.ColumnType_STRING, Description: "Specifies the ID of the current app release.", Transform: transform.FromField("CurrentRelease.Id")},
			{Name: "organization", Type: proto.ColumnType_JSON, Description: "Specifies the organization details where the app is deployed."},
			{Name: "autoscaling", Type: proto.ColumnType_JSON, Description: "Specifies the autoscaling information."},
			{Name: "config", Type: proto.ColumnType_JSON, Description: "Specifies the app configuration."},
			{Name: "health_checks", Type: proto.ColumnType_JSON, Description: "Specifies the app health check information."},
			{Name: "ip_addresses", Type: proto.ColumnType_JSON, Description: "A list of IP addresses associated with the app."},
		},
	}
}

//// LIST FUNCTION

func listFlyApps(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
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

	// Create client
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
