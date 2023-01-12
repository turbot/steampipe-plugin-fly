package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Machine struct {
	Config     MachineConfiguration `json:"config"`
	CreatedAt  string               `json:"createdAt"`
	Host       Host                 `json:"host"`
	ID         string               `json:"id"`
	InstanceID string               `json:"instanceId"`
	Name       string               `json:"name"`
	Region     string               `json:"region"`
	State      string               `json:"state"`
	UpdatedAt  string               `json:"updatedAt"`
}

type Mount struct {
	Path      string `json:"path"`
	Volume    string `json:"volume"`
	SizeGb    int    `json:"size_gb"`
	Encrypted bool   `json:"encrypted"`
}

type Host struct {
	ID string `json:"id"`
}

type MachineImageRef struct {
	Registry   string      `json:"registry"`
	Repository string      `json:"repository"`
	Tag        string      `json:"tag"`
	Digest     string      `json:"digest"`
	Labels     interface{} `json:"labels"`
}

type MachineConfiguration struct {
	Size     string          `json:"size"`
	Image    string          `json:"image"`
	Mounts   []Mount         `json:"mounts"`
	Env      interface{}     `json:"env"`
	ImageRef MachineImageRef `json:"image_ref"`
}

type ListMachinesRequestConfiguration struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned EndCursor value.
	Limit int

	// When paginating forwards, the cursor to continue.
	EndCursor string

	// The ID of the application.
	AppID string
	
	// The state of the machine.
	State string
}

type ListMachinesResponse struct {
	Machines Machines `json:"machines"`
}

type GetMachineResponse struct {
	Machine Machine `json:"machine"`
}

type Machines struct {
	Nodes      []Machine `json:"nodes"`
	PageInfo   PageInfo  `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

// __ListMachinesInput is used internally by genqlient
type __ListMachinesInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	AppID string `json:"app_id"`
	State string `json:"state"`
}

// __GetMachineInput is used internally by genqlient
type __GetMachineInput struct {
	MachineID string `json:"machineId"`
}

// Define the query
const (
	queryMachineList = `
query ListMachines($first: Int, $after: String, $appId: String, $state: String) {
  machines(first: $first, after: $after, appId: $appId, state: $state) {
    pageInfo {
      endCursor
      hasNextPage
    }
    totalCount
    nodes {
      name
      id
      state
      createdAt
      updatedAt
      instanceId
      host {
        id
      }
      region
      config
    }
  }
}
`

	queryMachineGet = `
query GetMachine($machineId: String!) {
  machine(machineId: $machineId) {
    name
    id
    state
    createdAt
    updatedAt
    instanceId
    host {
      id
    }
    region
    config
  }
}
`
)

// ListMachines returns all the machines
func ListMachines(
	ctx context.Context,
	client graphql.Client,
	options *ListMachinesRequestConfiguration,
) (*ListMachinesResponse, error) {

	// Check for options
	variables := &__ListMachinesInput{}
	if options.Limit > 0 {
		variables.First = options.Limit
	}

	if options.EndCursor != "" {
		variables.After = options.EndCursor
	}

	if options.AppID != "" {
		variables.AppID = options.AppID
	}

	if options.State != "" {
		variables.State = options.State
	}

	req := &graphql.Request{
		OpName:    "ListMachines",
		Query:     queryMachineList,
		Variables: variables,
	}
	var err error

	var data ListMachinesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetMachine returns the specified machine
func GetMachine(
	ctx context.Context,
	client graphql.Client,
	machineID string,
) (*GetMachineResponse, error) {
	req := &graphql.Request{
		OpName: "GetMachine",
		Query:  queryMachineGet,
		Variables: &__GetMachineInput{
			MachineID: machineID,
		},
	}
	var err error

	var data GetMachineResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
