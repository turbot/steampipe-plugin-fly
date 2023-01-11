package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Database struct {
	Name          string       `json:"name"`
	Id            string       `json:"id"`
	Hostname      string       `json:"hostname"`
	PublicUrl     string       `json:"publicUrl"`
	PrimaryRegion string       `json:"primaryRegion"`
	PrivateIp     string       `json:"privateIp"`
	Password      string       `json:"password"`
	AddOnPlanName string       `json:"addOnPlanName"`
	Options       interface{}  `json:"options"`
	Organization  Organization `json:"organization"`
	ReadRegions   []string     `json:"readRegions"`
	AddOnPlan     AddOnPlan    `json:"addOnPlan"`
}

type RedisDatabases struct {
	Nodes      []Database `json:"nodes"`
	PageInfo   PageInfo   `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

// ListRedisDatabasesResponse is returned by ListRedisDatabases on success.
type ListRedisDatabasesResponse struct {
	RedisDatabases RedisDatabases `json:"addOns"`
}

type GetRedisDatabaseResponse struct {
	RedisDatabase Database `json:"addOn"`
}

type ListRedisDatabasesRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string
}

// __ListRedisDatabasesInput is used internally by genqlient
type __ListRedisDatabasesInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	Type  string `json:"type"`
}

// __GetRedisDatabaseInput is used internally by genqlient
type __GetRedisDatabaseInput struct {
	Id string `json:"id"`
}

// Define the query
const (
	queryRedisDatabaseList = `
query ListRedisDatabases($type: AddOnType, $first: Int, $after: String) {
  addOns(type: $type, first: $first, after: $after) {
    pageInfo {
      hasNextPage
      endCursor
    }
    totalCount
    nodes {
      name
      id
      password
      primaryRegion
      privateIp
      publicUrl
      hostname
      options
      organization {
        id
      }
      readRegions
      addOnPlanName
      addOnPlan {
        displayName
        id
        name
        maxCommandsPerSec
        maxConcurrentConnections
        maxDailyCommands
        maxDailyBandwidth
        maxDataSize
        maxRequestSize
        pricePerMonth
      }
    }
  }
}
`

	queryRedisDatabaseGet = `
query GetRedisDatabase($id: ID) {
  addOn(id: $id) {
    name
    id
    password
    primaryRegion
    privateIp
    publicUrl
    hostname
    options
    organization {
      id
    }
    readRegions
    addOnPlanName
    addOnPlan {
      displayName
      id
      name
      maxCommandsPerSec
      maxConcurrentConnections
      maxDailyCommands
      maxDailyBandwidth
      maxDataSize
      maxRequestSize
      pricePerMonth
    }
  }
}
`
)

// ListRedisDatabases returns all the organizations the user has access to
func ListRedisDatabases(
	ctx context.Context,
	client graphql.Client,
	options *ListRedisDatabasesRequestConfiguration,
) (*ListRedisDatabasesResponse, error) {

	// Check for options
	variables := &__ListRedisDatabasesInput{
		// Should be always "redis"
		Type: "redis",
	}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	req := &graphql.Request{
		OpName:    "ListRedisDatabases",
		Query:     queryRedisDatabaseList,
		Variables: variables,
	}
	var err error

	var data ListRedisDatabasesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetRedisDatabase returns the specified organization
func GetRedisDatabase(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*GetRedisDatabaseResponse, error) {
	req := &graphql.Request{
		OpName: "GetRedisDatabase",
		Query:  queryRedisDatabaseGet,
		Variables: &__GetRedisDatabaseInput{
			Id: id,
		},
	}
	var err error

	var data GetRedisDatabaseResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
