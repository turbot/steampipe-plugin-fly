package apiClient

import (
	"context"
	"net/http"
	"time"

	"github.com/turbot/steampipe-plugin-fly/utils"

	"github.com/Khan/genqlient/graphql"
)

// Fly API Client
type Client struct {
	Token   *string
	Graphql graphql.Client
}

func CreateClient(ctx context.Context, config ClientConfig) (*Client, error) {
	h := http.Client{Timeout: 60 * time.Second, Transport: &utils.Transport{UnderlyingTransport: http.DefaultTransport, Token: *config.Token, Ctx: ctx}}

	return &Client{
		Token:   config.Token,
		Graphql: graphql.NewClient("https://api.fly.io/graphql", &h),
	}, nil
}
