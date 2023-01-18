package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Organization struct {
	ActiveDiscountName     string `json:"activeDiscountName"`
	AddOnSSOLink           string `json:"addOnSsoLink"`
	BillingStatus          string `json:"billingStatus"`
	CreditBalance          int    `json:"creditBalance"`
	CreditBalanceFormatted string `json:"creditBalanceFormatted"`
	ID                     string `json:"id"`
	InternalNumericID      string `json:"internalNumericId"`
	IsCreditCardSaved      bool   `json:"isCreditCardSaved"`
	Name                   string `json:"name"`
	RemoteBuilderImage     string `json:"remoteBuilderImage"`
	SSHCertificate         string `json:"sshCertificate"`
	Slug                   string `json:"slug"`
	Trust                  string `json:"trust"`
	Type                   string `json:"type"`
	ViewerRole             string `json:"viewerRole"`
}

type ListOrganizationsRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string
}

// ListOrganizationsResponse is returned by ListOrganizations on success.
type ListOrganizationsResponse struct {
	Organizations Organizations `json:"organizations"`
}

type GetOrganizationResponse struct {
	Organization Organization `json:"organization"`
}

type Organizations struct {
	Nodes      []Organization `json:"nodes"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

// __ListOrganizationsInput is used internally by genqlient
type __ListOrganizationsInput struct {
	First int    `json:"first"`
	After string `json:"after"`
}

// __GetOrganizationInput is used internally by genqlient
type __GetOrganizationInput struct {
	Slug string `json:"slug"`
}

// Define the query
const (
	queryOrganizationList = `
query ListOrganizations($first: Int, $after: String) {
  organizations(first: $first, after: $after) {
    pageInfo {
      hasNextPage
      endCursor
    }
    totalCount
    nodes {
      name
      id
      slug
      billingStatus
      creditBalance
      creditBalanceFormatted
      sshCertificate
      internalNumericId
      isCreditCardSaved
      activeDiscountName
      type
      viewerRole
      remoteBuilderImage
      addOnSsoLink
      trust
    }
  }
}
`

	queryOrganizationGet = `
query GetOrganization($slug: String) {
  organization(slug: $slug) {
    name
    id
    slug
    billingStatus
    creditBalance
    creditBalanceFormatted
    sshCertificate
    internalNumericId
    isCreditCardSaved
    activeDiscountName
    type
    viewerRole
    remoteBuilderImage
    addOnSsoLink
    trust
  }
}
`
)

// ListOrganizations returns all the organizations the user has access to
func ListOrganizations(
	ctx context.Context,
	client graphql.Client,
	options *ListOrganizationsRequestConfiguration,
) (*ListOrganizationsResponse, error) {

	// Check for options
	variables := &__ListOrganizationsInput{}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	req := &graphql.Request{
		OpName:    "ListOrganizations",
		Query:     queryOrganizationList,
		Variables: variables,
	}
	var err error

	var data ListOrganizationsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetOrganization returns the specified organization
func GetOrganization(
	ctx context.Context,
	client graphql.Client,
	slug string,
) (*GetOrganizationResponse, error) {
	req := &graphql.Request{
		OpName: "GetOrganization",
		Query:  queryOrganizationGet,
		Variables: &__GetOrganizationInput{
			Slug: slug,
		},
	}
	var err error

	var data GetOrganizationResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
