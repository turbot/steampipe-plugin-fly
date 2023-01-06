package apiClient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

type ListAppsRequestConfiguration struct {
	Limit     int
	EndCursor string
}

// ListAppsResponse is returned by ListApps on success.
type ListAppsResponse struct {
	Apps Apps `json:"apps"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

type Apps struct {
	Nodes      []provider.GetFullAppApp `json:"nodes"`
	PageInfo   PageInfo                 `json:"pageInfo"`
	TotalCount int                      `json:"totalCount"`
}

// __GetFullAppInput is used internally by genqlient
type __GetFullAppInput struct {
	Name string `json:"name"`
}

type __ListAppsInput struct {
	First int    `json:"first"`
	After string `json:"after"`
}

const (
	queryAppList = `
query ListApps($first: Int, $after: String) {
	apps(first: $first, after: $after) {
		pageInfo {
      hasNextPage
			endCursor
    }
    totalCount
		nodes {
			name
			network
			organization {
				id
				slug
			}
			autoscaling {
				preferredRegion
				regions {
					code
				}
			}
			appUrl
			hostname
			id
			status
			deployed
			currentRelease {
				id
			}
			config {
				definition
			}
			healthChecks {
				nodes {
					name
					status
				}
			}
			ipAddresses {
				nodes {
					address
					id
				}
			}
			role {
				__typename
				name
			}
		}
	}
}
`

	queryAppGet = `
query GetFullApp ($name: String) {
	app(name: $name) {
		name
		network
		organization {
			id
			slug
		}
		autoscaling {
			preferredRegion
			regions {
				code
			}
		}
		appUrl
		hostname
		id
		status
		deployed
		currentRelease {
			id
		}
		config {
			definition
		}
		healthChecks {
			nodes {
				name
				status
			}
		}
		ipAddresses {
			nodes {
				address
				id
			}
		}
		role {
			__typename
			name
		}
	}
}
`
)

// List apps
func ListApps(
	ctx context.Context,
	client graphql.Client,
	options *ListAppsRequestConfiguration,
) (*ListAppsResponse, error) {
	// Check for options
	variables := &__ListAppsInput{}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	// Create request
	req := &graphql.Request{
		OpName:    "ListApps",
		Query:     queryAppList,
		Variables: variables,
	}
	var err error

	var data ListAppsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// Get app
func GetFullApp(
	ctx context.Context,
	client graphql.Client,
	name string,
) (*provider.GetFullAppResponse, error) {
	req := &graphql.Request{
		OpName: "GetFullApp",
		Query:  queryAppGet,
		Variables: &__GetFullAppInput{
			Name: name,
		},
	}
	var err error

	var data provider.GetFullAppResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
