package apiClient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

type ListAppsRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string
}

// ListAppsResponse is returned by ListApps on success.
type ListAppsResponse struct {
	Apps Apps `json:"apps"`
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

// __ListAppsInput is used internally by genqlient
type __ListAppsInput struct {
	First int    `json:"first"`
	After string `json:"after"`
}

// Define the query
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

// ListApps returns all the fly apps resource
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

// GetFullApp returns the specified the fly app resource
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
