package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
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
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the machine."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the machine.", Transform: transform.FromGo()},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The current status of the machine."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The region where the machine is created."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the machine was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the machine was last modified."},
			{Name: "instance_id", Type: proto.ColumnType_STRING, Description: "Specifies the instance ID.", Transform: transform.FromField("InstanceID")},
			{Name: "host_id", Type: proto.ColumnType_STRING, Description: "Specifies the machine host ID.", Transform: transform.FromField("Host.ID")},
			{Name: "config", Type: proto.ColumnType_JSON, Description: "Specifies the machine configuration."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "Specifies the application.", Transform: transform.FromQual("app_id")},
		},
	}
}

//// LIST FUNCTION

func listFlyMachines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_machine.listFlyMachines", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListMachinesRequestConfiguration{}

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
		query, err := flyapi.ListMachines(context.Background(), conn.Graphql, options)
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

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_machine.getFlyMachine", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetMachine(context.Background(), conn.Graphql, machineID)
	if err != nil {
		plugin.Logger(ctx).Error("fly_machine.getFlyMachine", "query_error", err)
		return nil, err
	}

	return query.Machine, nil
}
