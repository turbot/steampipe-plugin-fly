package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
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
			ParentHydrate: listFlyOrganizations,
			Hydrate:       listFlyOrganizationMembers,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the member.", Transform: transform.FromField("Member.Name")},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the member.", Transform: transform.FromField("Member.Id")},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address of the member.", Transform: transform.FromField("Member.Email")},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "The username of the member.", Transform: transform.FromField("Member.Username")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the member was created.", Transform: transform.FromField("Member.CreatedAt")},
			{Name: "two_factor_protection", Type: proto.ColumnType_BOOL, Description: "If true, two-factor authentication is enabled for the member.", Transform: transform.FromField("Member.TwoFactorProtection")},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The role of the member in the organization."},
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "The avatar URL of the member.", Transform: transform.FromField("Member.AvatarUrl")},
			{Name: "last_region", Type: proto.ColumnType_STRING, Description: "Specifies the region the member is recently accessed to.", Transform: transform.FromField("Member.LastRegion")},
			{Name: "has_node_proxy_apps", Type: proto.ColumnType_BOOL, Description: "True, if the member has node proxy app.", Transform: transform.FromField("Member.HasNodeProxyApps")},
			{Name: "trust", Type: proto.ColumnType_STRING, Description: "Specifies the trust level. Possible values are: UNKNOWN, RESTRICTED, BANNED, LOW, HIGH.", Transform: transform.FromField("Member.Trust")},
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "The ID of the organization.", Transform: transform.FromField("Member.OrganizationId")},
			{Name: "feature_flags", Type: proto.ColumnType_JSON, Description: "A list of feature flags.", Transform: transform.FromField("Member.FeatureFlags")},
		},
	}
}

//// LIST FUNCTION

func listFlyOrganizationMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	orgData := h.Item.(flyapi.Organization)
	orgID := orgData.ID

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_organization_member.listFlyOrganizationMembers", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListOrgMembersRequestConfiguration{
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
		query, err := flyapi.ListOrganizationMembers(context.Background(), conn.Graphql, options)
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
