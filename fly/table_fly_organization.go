package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFlyOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_organization",
		Description: "Fly Organization",
		List: &plugin.ListConfig{
			Hydrate: listFlyOrganizations,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyOrganization,
			KeyColumns: plugin.SingleColumn("slug"),
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the organization."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the organization.", Transform: transform.FromGo()},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The organization slug name."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the organization."},
			{Name: "billing_status", Type: proto.ColumnType_STRING, Description: "The billing status of the organization."},
			{Name: "credit_balance", Type: proto.ColumnType_INT, Description: "The current remaining credit balance of the organization."},
			{Name: "credit_balance_formatted", Type: proto.ColumnType_STRING, Description: "The formatted current remaining credit balance of the organization."},
			{Name: "is_credit_card_saved", Type: proto.ColumnType_BOOL, Description: "If true, a valid credit card is provided for billing."},
			{Name: "ssh_certificate", Type: proto.ColumnType_STRING, Description: "Specifies the SSH certificate.", Transform: transform.FromField("SSHCertificate")},
			{Name: "internal_numeric_id", Type: proto.ColumnType_STRING, Description: "The internal numeric ID of the organization.", Transform: transform.FromField("InternalNumericID")},
			{Name: "active_discount_name", Type: proto.ColumnType_STRING, Description: "Specifies the active discount name."},
			{Name: "viewer_role", Type: proto.ColumnType_STRING, Description: "Indicaßßßßtes who can view the details."},
			{Name: "remote_builder_image", Type: proto.ColumnType_STRING, Description: "Specifies the remote builder image of the organization."},
			{Name: "add_on_sso_link", Type: proto.ColumnType_STRING, Description: "Specifies the addOn SSO link.", Transform: transform.FromField("AddOnSSOLink")},
			{Name: "trust", Type: proto.ColumnType_STRING, Description: "Specifies the trust level. Possible values are: UNKNOWN, RESTRICTED, BANNED, LOW, HIGH."},
		},
	}
}

//// LIST FUNCTION

func listFlyOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization.listFlyOrganizations", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListOrganizationsRequestConfiguration{}

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
		query, err := flyapi.ListOrganizations(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_organization.listFlyOrganizations", "query_error", err)
			return nil, err
		}

		for _, org := range query.Organizations.Nodes {
			d.StreamListItem(ctx, org)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.Organizations.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.Organizations.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	slug := d.EqualsQualString("slug")

	// Return nil, if empty
	if slug == "" {
		return nil, nil
	}

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization.getFlyOrganization", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetOrganization(context.Background(), conn.Graphql, slug)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization.getFlyOrganization", "query_error", err)
		return nil, err
	}

	return query.Organization, nil
}
