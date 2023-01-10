package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/apiClient"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFlyOrganizationMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_organization_member",
		Description: "Fly Organization Member",
		List: &plugin.ListConfig{
			ParentHydrate: listFlyOrganization,
			Hydrate:       listFlyOrganizationMembers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.Name"),
			},
			{
				Name:        "id",
				Description: "A unique identifier of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.Id"),
			},
			{
				Name:        "email",
				Description: "The email address of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.Email"),
			},
			{
				Name:        "username",
				Description: "The username of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.Username"),
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the member was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Member.CreatedAt"),
			},
			{
				Name:        "two_factor_protection",
				Description: "If true, two-factor authentication is enabled for the member.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Member.TwoFactorProtection"),
			},
			{
				Name:        "role",
				Description: "The role of the member in the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "avatar_url",
				Description: "The avatar URL of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.AvatarUrl"),
			},
			{
				Name:        "last_region",
				Description: "Specifies the region the member is recently accessed to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.LastRegion"),
			},
			{
				Name:        "has_node_proxy_apps",
				Description: "",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Member.HasNodeProxyApps"),
			},
			{
				Name:        "trust",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.Trust"),
			},
			{
				Name:        "organization_id",
				Description: "The ID of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Member.OrganizationId"),
			},
			{
				Name:        "feature_flags",
				Description: "A list of feature flags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Member.FeatureFlags"),
			},
		},
	}
}

//// LIST FUNCTION

func listFlyOrganizationMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	orgData := h.Item.(apiClient.Organization)
	orgID := orgData.ID

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization_member.listFlyOrganizationMembers", "connection_error", err)
		return nil, err
	}

	options := &apiClient.ListOrgMembersRequestConfiguration{
		OrgId: orgID,
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
		query, err := apiClient.ListOrganizationMembers(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_organization_member.listFlyOrganizationMembers", "query_error", err)
			return nil, err
		}

		for _, e := range query.Organization.Members.Edges {
			e.Member.OrganizationId = orgID
			d.StreamListItem(ctx, e)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.Organization.Members.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.Organization.Members.PageInfo.EndCursor
	}

	return nil, nil
}
