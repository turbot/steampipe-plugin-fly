package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Location struct {
	Name        string
	Title       string
	State       string
	Locality    string
	Country     string
	Coordinates []float64
}

// ListLocationsResponse is returned by ListLocations on success.
type ListLocationsResponse struct {
	Locations []Location `json:"checkLocations"`
}

// Define the query
const (
	queryLocationList = `
query ListLocations {
	checkLocations {
		name
		title
		state
		locality
		country
		coordinates
	}
}
`
)

// ListLocations returns all the available locations
func ListLocations(
	ctx context.Context,
	client graphql.Client,
) (*ListLocationsResponse, error) {

	req := &graphql.Request{
		OpName: "ListLocations",
		Query:  queryLocationList,
	}
	var err error

	var data ListLocationsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
