package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

//// TABLE DEFINITION

func tableFlyIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_ip",
		Description: "Fly IP Address",
		List: &plugin.ListConfig{
			ParentHydrate: listFlyApps,
			Hydrate:       listFlyIPs,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyIP,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "address", Type: proto.ColumnType_IPADDR, Description: "The IP address"},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the IP address."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the IP address was created."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region where the IP address is created."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Specifies the type of the IP address."},
		},
	}
}

//// LIST FUNCTION

func listFlyIPs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	appData := h.Item.(provider.GetFullAppApp)
	appID := appData.Name

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_ip.listFlyIPs", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListIPAddressesRequestConfiguration{
		AppId: appID,
	}

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
		query, err := flyapi.ListIPAddresses(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_ip.listFlyIPs", "query_error", err)
			return nil, err
		}

		for _, ip := range query.App.IPAddresses.Nodes {
			d.StreamListItem(ctx, ip)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.App.IPAddresses.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.App.IPAddresses.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")
	if id == "" {
		return nil, nil
	}

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_ip.getFlyIP", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetIPAddress(context.Background(), conn.Graphql, id)
	if err != nil {
		plugin.Logger(ctx).Error("fly_ip.getFlyIP", "query_error", err)
		return nil, err
	}

	return query.IPAddress, nil
}
