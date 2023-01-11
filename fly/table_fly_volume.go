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
			ParentHydrate: listFlyApp,
			Hydrate:       listFlyVolumes,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyVolume,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier of the volume.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "state",
				Description: "The configuration state of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "encrypted",
				Description: "If true, the volume is encrypted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "region",
				Description: "The region where the volume is created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "size_gb",
				Description: "Specifies the size of the volume.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the volume was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "status",
				Description: "The current status of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "internal_id",
				Description: "Specifies the internal ID of the volume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "used_bytes",
				Description: "Specifies the amount of storage used in bytes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "host_id",
				Description: "Specifies the volume host ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Host.ID"),
			},
			{
				Name:        "attached_machine_id",
				Description: "Specifies the ID of the machine; the volume is attached.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AttachedMachine.ID"),
			},
			{
				Name:        "app_id",
				Description: "Specifies the ID of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("App.Id"),
			},
		},
	}
}

//// LIST FUNCTION

func listFlyVolumes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	appData := h.Item.(provider.GetFullAppApp)
	appID := appData.Name

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
	if volumeID == "" {
		return nil, nil
	}

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
