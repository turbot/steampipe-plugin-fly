package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

type Volume struct {
	Name            string                 `json:"name"`
	ID              string                 `json:"id"`
	SizeGb          int                    `json:"sizeGb"`
	State           string                 `json:"state"`
	Status          string                 `json:"status"`
	Region          string                 `json:"region"`
	CreatedAt       string                 `json:"createdAt"`
	InternalId      string                 `json:"internalId"`
	UsedBytes       string                 `json:"usedBytes"`
	Encrypted       bool                   `json:"encrypted"`
	Host            Host                   `json:"host"`
	AttachedMachine Machine                `json:"attachedMachine"`
	App             provider.GetFullAppApp `json:"app"`
}

type Volumes struct {
	Nodes      []Volume `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount"`
}

type VolumeQueryApp struct {
	Volumes Volumes `json:"volumes"`
}

type ListVolumesResponse struct {
	App VolumeQueryApp `json:"app"`
}

type GetVolumeResponse struct {
	Volume Volume `json:"volume"`
}

type ListVolumesRequestConfiguration struct {
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

// __ListVolumesInput is used internally by genqlient
type __ListVolumesInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	AppId string `json:"appId"`
}

// __GetVolumeInput is used internally by genqlient
type __GetVolumeInput struct {
	Id string `json:"id"`
}

// Define the query
const (
	queryVolumeList = `
query ListVolumes($appId: String, $first: Int, $after: String) {
  app(name: $appId) {
    volumes(first: $first, after: $after) {
      pageInfo {
        endCursor
        hasNextPage
      }
      totalCount
      nodes {
        name
        region
        id
        state
        sizeGb
        status
        createdAt
        internalId
        usedBytes
        encrypted
        host {
          id
        }
        attachedMachine {
          id
        }
        app {
          id
        }
      }
    }
  }
}
`

	queryVolumeGet = `
query GetVolume($id: ID!) {
  volume(id: $id) {
    name
    region
    id
    state
    sizeGb
    status
    createdAt
    internalId
    usedBytes
    encrypted
    host {
      id
    }
    attachedMachine {
      id
    }
    app {
      id
    }
  }
}
`
)

// ListVolumes returns all the volumes
func ListVolumes(
	ctx context.Context,
	client graphql.Client,
	options *ListVolumesRequestConfiguration,
) (*ListVolumesResponse, error) {

	// Check for options
	variables := &__ListVolumesInput{}
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
		OpName:    "ListVolumes",
		Query:     queryVolumeList,
		Variables: variables,
	}
	var err error

	var data ListVolumesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetVolume returns the specified volume
func GetVolume(
	ctx context.Context,
	client graphql.Client,
	Id string,
) (*GetVolumeResponse, error) {
	req := &graphql.Request{
		OpName: "GetVolume",
		Query:  queryVolumeGet,
		Variables: &__GetVolumeInput{
			Id: Id,
		},
	}
	var err error

	var data GetVolumeResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
