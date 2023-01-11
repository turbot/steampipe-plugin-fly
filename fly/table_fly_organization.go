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
			Hydrate: listFlyOrganization,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyOrganization,
			KeyColumns: plugin.SingleColumn("slug"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "slug",
				Description: "The organization slug name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "billing_status",
				Description: "The billing status of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "credit_balance",
				Description: "The current remaining credit balance of the organization.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "credit_balance_formatted",
				Description: "The formatted current remaining credit balance of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_credit_card_saved",
				Description: "If true, a valid credit card is provided for billing.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ssh_certificate",
				Description: "Specifies the SSH certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SSHCertificate"),
			},
			{
				Name:        "internal_numeric_id",
				Description: "The internal numeric ID of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InternalNumericID"),
			},
			{
				Name:        "active_discount_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "viewer_role",
				Description: "Indicates who can view the details.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remote_builder_image",
				Description: "Specifies the remote builder image of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "add_on_sso_link",
				Description: "Specifies the addOn SSO link.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AddOnSSOLink"),
			},
			{
				Name:        "trust",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listFlyOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization.listFlyOrganization", "connection_error", err)
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
			plugin.Logger(ctx).Error("fly_organization.listFlyOrganization", "query_error", err)
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
	if slug == "" {
		return nil, nil
	}

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
