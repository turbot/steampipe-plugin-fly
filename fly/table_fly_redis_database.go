package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFlyRedisDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_redis_database",
		Description: "Fly Redis Database",
		List: &plugin.ListConfig{
			Hydrate: listFlyRedisDatabases,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyRedisDatabase,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the database."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the database."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The database hostname."},
			{Name: "public_url", Type: proto.ColumnType_STRING, Description: "The public URL of the database."},
			{Name: "primary_region", Type: proto.ColumnType_STRING, Description: "The primary region where the database is located."},
			{Name: "private_ip", Type: proto.ColumnType_IPADDR, Description: "Specifies the private IP address of the database."},
			{Name: "password", Type: proto.ColumnType_STRING, Description: "The database password."},
			{Name: "add_on_plan_name", Type: proto.ColumnType_STRING, Description: "Specifies the database plan."},
			{Name: "options", Type: proto.ColumnType_JSON, Description: "The database options."},
			{Name: "read_regions", Type: proto.ColumnType_JSON, Description: "A list of database replica regions."},
			{Name: "add_on_plan", Type: proto.ColumnType_JSON, Description: "Specifies the add-on plan."},
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "Specifies the organization.", Transform: transform.FromField("Organization.ID")},
		},
	}
}

//// LIST FUNCTION

func listFlyRedisDatabases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_redis_database.listFlyRedisDatabases", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListRedisDatabasesRequestConfiguration{}

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
		query, err := flyapi.ListRedisDatabases(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_redis_database.listFlyRedisDatabases", "query_error", err)
			return nil, err
		}

		for _, db := range query.RedisDatabases.Nodes {
			d.StreamListItem(ctx, db)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.RedisDatabases.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.RedisDatabases.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyRedisDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Return nil, if empty
	if id == "" {
		return nil, nil
	}

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_redis_database.getFlyRedisDatabase", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetRedisDatabase(context.Background(), conn.Graphql, id)
	if err != nil {
		plugin.Logger(ctx).Error("fly_redis_database.getFlyRedisDatabase", "query_error", err)
		return nil, err
	}

	return query.RedisDatabase, nil
}
