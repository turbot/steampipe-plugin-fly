package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/apiClient"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFlyMachine(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_machine",
		Description: "Fly Machine",
		List: &plugin.ListConfig{
			Hydrate: listFlyMachines,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "state", Require: plugin.Optional},
				{Name: "app_id", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyMachine,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the machine.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier of the machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "state",
				Description: "The current status of the machine.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "The region where the machine is created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the machine was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "updated_at",
				Description: "The timestamp when the machine was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "instance_id",
				Description: "Specifies the instance ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceID"),
			},
			{
				Name:        "host_id",
				Description: "Specifies the machine host ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Host.ID"),
			},
			{
				Name:        "config",
				Description: "Specifies the machine configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "app_id",
				Description: "Specifies the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("app_id"),
			},
		},
	}
}

//// LIST FUNCTION

func listFlyMachines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_machine.listFlyMachines", "connection_error", err)
		return nil, err
	}

	options := &apiClient.ListMachinesRequestConfiguration{}

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

	// Check for filters
	if d.EqualsQualString("state") != "" {
		options.State = d.EqualsQualString("state")
	}

	if d.EqualsQualString("app_id") != "" {
		options.AppID = d.EqualsQualString("app_id")
	}

	for {
		query, err := apiClient.ListMachines(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_machine.listFlyMachines", "query_error", err)
			return nil, err
		}

		for _, machine := range query.Machines.Nodes {
			d.StreamListItem(ctx, machine)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.Machines.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.Machines.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyMachine(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	machineID := d.EqualsQualString("id")
	if machineID == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_machine.getFlyMachine", "connection_error", err)
		return nil, err
	}

	query, err := apiClient.GetMachine(context.Background(), conn.Graphql, machineID)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization.getFlyOrganization", "query_error", err)
		return nil, err
	}

	return query.Machine, nil
}
