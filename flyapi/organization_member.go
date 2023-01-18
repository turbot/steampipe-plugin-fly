package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Member struct {
	AvatarUrl           string   `json:"avatarUrl"`
	CreatedAt           string   `json:"createdAt"`
	Email               string   `json:"email"`
	FeatureFlags        []string `json:"featureFlags"`
	HasNodeProxyApps    bool     `json:"hasNodeproxyApps"`
	Id                  string   `json:"id"`
	LastRegion          string   `json:"lastRegion"`
	Name                string   `json:"name"`
	OrganizationId      string   `json:"-"`
	Trust               string   `json:"trust"`
	TwoFactorProtection bool     `json:"twoFactorProtection"`
	Username            string   `json:"username"`
}

type OrganizationMembershipsEdge struct {
	Member Member `json:"node"`
	Role   string `json:"role"`
}

type Members struct {
	Edges      []OrganizationMembershipsEdge `json:"edges"`
	PageInfo   PageInfo                      `json:"pageInfo"`
	TotalCount int                           `json:"totalCount"`
}

type MemberQueryOrganization struct {
	Members Members `json:"members"`
}

type ListOrgMembersResponse struct {
	Organization MemberQueryOrganization `json:"organization"`
}

type ListOrgMembersRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string

	// The ID of the organization.
	//
	// Required
	OrgId string
}

// __ListOrgMembersInput is used internally by genqlient
type __ListOrgMembersInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	OrgId string `json:"orgID"`
}

// Define the query
const (
	queryOrgMemberList = `
query ListOrganizationMembers($orgID: ID, $first: Int, $after: String) {
  organization(id: $orgID) {
    members(first: $first, after: $after) {
      pageInfo {
        hasNextPage
        endCursor
      }
      totalCount
      edges {
        node {
          username
          avatarUrl
          name
          email
          createdAt
          lastRegion
          twoFactorProtection
          hasNodeproxyApps
          featureFlags
          trust
          id
        }
        role
      }
    }
  }
}
`
)

// ListOrganizationMembers returns all the members of an organization
func ListOrganizationMembers(
	ctx context.Context,
	client graphql.Client,
	options *ListOrgMembersRequestConfiguration,
) (*ListOrgMembersResponse, error) {

	// Check for options
	variables := &__ListOrgMembersInput{}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	if options.OrgId != "" {
		variables.OrgId = options.OrgId
	}

	req := &graphql.Request{
		OpName:    "ListOrganizationMembers",
		Query:     queryOrgMemberList,
		Variables: variables,
	}
	var err error

	var data ListOrgMembersResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
