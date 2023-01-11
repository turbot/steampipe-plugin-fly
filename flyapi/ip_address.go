package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type IPAddress struct {
	Id        string `json:"id"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	Region    string `json:"region"`
	Type      string `json:"type"`
}

type IPAddresses struct {
	Nodes      []IPAddress `json:"nodes"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type IPAddressQueryApp struct {
	IPAddresses IPAddresses `json:"ipAddresses"`
}

type ListIPAddressesResponse struct {
	App IPAddressQueryApp `json:"app"`
}

type GetIPAddressResponse struct {
	IPAddress IPAddress `json:"ipAddress"`
}

type ListIPAddressesRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string

	// The ID of the application.
	//
	// Required
	AppId string
}

// __ListIPAddressesInput is used internally by genqlient
type __ListIPAddressesInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	AppId string `json:"appId"`
}

// __GetIPAddressInput is used internally by genqlient
type __GetIPAddressInput struct {
	Id string `json:"id"`
}

// Define the query
const (
	queryIPAddressList = `
query ListIPAddresses($appId: String, $first: Int, $after: String) {
  app(name: $appId) {
    ipAddresses(first: $first, after: $after) {
      pageInfo {
        endCursor
        hasNextPage
      }
      totalCount
      nodes {
        address
        createdAt
        id
        region
        type
      }
    }
  }
}
`

	queryIPAddressGet = `
query GetIPAddress($id: ID!) {
  ipAddress(id: $id) {
    address
    createdAt
    id
    region
    type
  }
}
`
)

// ListIPAddresses returns all the volumes
func ListIPAddresses(
	ctx context.Context,
	client graphql.Client,
	options *ListIPAddressesRequestConfiguration,
) (*ListIPAddressesResponse, error) {

	// Check for options
	variables := &__ListIPAddressesInput{}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	if options.AppId != "" {
		variables.AppId = options.AppId
	}

	req := &graphql.Request{
		OpName:    "ListIPAddresses",
		Query:     queryIPAddressList,
		Variables: variables,
	}
	var err error

	var data ListIPAddressesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetIPAddress returns the specified volume
func GetIPAddress(
	ctx context.Context,
	client graphql.Client,
	Id string,
) (*GetIPAddressResponse, error) {
	req := &graphql.Request{
		OpName: "GetIPAddress",
		Query:  queryIPAddressGet,
		Variables: &__GetIPAddressInput{
			Id: Id,
		},
	}
	var err error

	var data GetIPAddressResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
