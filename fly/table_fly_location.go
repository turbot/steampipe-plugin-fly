package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableFlyLocation(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_location",
		Description: "Fly Location",
		List: &plugin.ListConfig{
			Hydrate: listFlyLocations,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the location."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Specifies the title."},
			{Name: "locality", Type: proto.ColumnType_STRING, Description: "Specifies the locality."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The state of the location."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Specifies the country name."},
			{Name: "coordinates", Type: proto.ColumnType_JSON, Description: "Specifies the latitude and longitude of the location."},
		},
	}
}

//// LIST FUNCTION

func listFlyLocations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_location.listFlyLocations", "connection_error", err)
		return nil, err
	}
	// As of Jan 13, 2023, paging is not supported
	query, err := flyapi.ListLocations(context.Background(), conn.Graphql)
	if err != nil {
		plugin.Logger(ctx).Error("fly_location.listFlyLocations", "query_error", err)
		return nil, err
	}

	for _, org := range query.Locations {
		d.StreamListItem(ctx, org)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
