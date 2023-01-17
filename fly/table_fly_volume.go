package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

//// TABLE DEFINITION

func tableFlyVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_volume",
		Description: "Fly Volume",
		List: &plugin.ListConfig{
			ParentHydrate: listFlyApps,
			Hydrate:       listFlyVolumes,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyVolume,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the volume."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the volume.", Transform: transform.FromGo()},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The configuration state of the volume."},
			{Name: "encrypted", Type: proto.ColumnType_BOOL, Description: "If true, the volume is encrypted."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region where the volume is created."},
			{Name: "size_gb", Type: proto.ColumnType_INT, Description: "Specifies the size of the volume."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the volume was created."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the volume."},
			{Name: "internal_id", Type: proto.ColumnType_STRING, Description: "Specifies the internal ID of the volume."},
			{Name: "used_bytes", Type: proto.ColumnType_STRING, Description: "Specifies the amount of storage used in bytes."},
			{Name: "host_id", Type: proto.ColumnType_STRING, Description: "Specifies the volume host ID.", Transform: transform.FromField("Host.ID")},
			{Name: "attached_machine_id", Type: proto.ColumnType_STRING, Description: "Specifies the ID of the machine; the volume is attached.", Transform: transform.FromField("AttachedMachine.ID")},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "Specifies the ID of the application.", Transform: transform.FromField("App.Id")},
		},
	}
}

//// LIST FUNCTION

func listFlyVolumes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	appData := h.Item.(provider.GetFullAppApp)
	appID := appData.Name

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_volume.listFlyVolumes", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListVolumesRequestConfiguration{
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
		query, err := flyapi.ListVolumes(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_volume.listFlyVolumes", "query_error", err)
			return nil, err
		}

		for _, volume := range query.App.Volumes.Nodes {
			d.StreamListItem(ctx, volume)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.App.Volumes.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.App.Volumes.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	volumeID := d.EqualsQualString("id")

	// Return nil, if empty
	if volumeID == "" {
		return nil, nil
	}

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_volume.getFlyVolume", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetVolume(context.Background(), conn.Graphql, volumeID)
	if err != nil {
		plugin.Logger(ctx).Error("fly_volume.getFlyVolume", "query_error", err)
		return nil, err
	}

	return query.Volume, nil
}
