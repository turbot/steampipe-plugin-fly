package flyapi

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type AppCertificate struct {
	CertificateAuthority      string `json:"certificateAuthority"`
	ClientStatus              string `json:"clientStatus"`
	CreatedAt                 string `json:"createdAt"`
	DnsProvider               string `json:"dnsProvider"`
	DnsValidationHostname     string `json:"dnsValidationHostname"`
	DnsValidationInstructions string `json:"dnsValidationInstructions"`
	DnsValidationTarget       string `json:"dnsValidationTarget"`
	Domain                    string `json:"domain"`
	Hostname                  string `json:"hostname"`
	Id                        string `json:"id"`
	IsAcmeAlpnConfigured      bool   `json:"isAcmeAlpnConfigured"`
	IsAcmeDnsConfigured       bool   `json:"isAcmeDnsConfigured"`
	IsApex                    bool   `json:"isApex"`
	IsConfigured              bool   `json:"isConfigured"`
	IsWildcard                bool   `json:"isWildcard"`
	Source                    string `json:"source"`
	Verified                  bool   `json:"check"`
}

type Certificates struct {
	Nodes      []AppCertificate `json:"nodes"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

type CertificateQueryApp struct {
	Certificates Certificates `json:"certificates"`
}

type ListAppCertificatesResponse struct {
	App CertificateQueryApp `json:"app"`
}

type GetAppCertificateResponse struct {
	Certificate AppCertificate `json:"certificate"`
}

type ListAppCertificatesRequestConfiguration struct {
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

// __ListAppCertificatesInput is used internally by genqlient
type __ListAppCertificatesInput struct {
	First int    `json:"first"`
	After string `json:"after"`
	AppId string `json:"appId"`
}

// __GetAppCertificateInput is used internally by genqlient
type __GetAppCertificateInput struct {
	Id string `json:"id"`
}

// Define the query
const (
	queryAppCertificateList = `
query ListAppCertificates($appId: String, $first: Int, $after: String) {
  app(name: $appId) {
    name
    certificates(first: $first, after: $after) {
      pageInfo {
        endCursor
        hasNextPage
      }
      totalCount
      nodes {
        hostname
        id
        domain
        source
        createdAt
        certificateAuthority
        clientStatus
        dnsProvider
        dnsValidationHostname
        dnsValidationInstructions
        dnsValidationTarget
        check
        isAcmeAlpnConfigured
        isAcmeDnsConfigured
        isApex
        isConfigured
        isWildcard
      }
    }
  }
}
`

	queryAppCertificateGet = `
query GetAppCertificate($id: ID!) {
  certificate(id: $id) {
    hostname
    id
    domain
    source
    createdAt
    certificateAuthority
    clientStatus
    dnsProvider
    dnsValidationHostname
    dnsValidationInstructions
    dnsValidationTarget
    check
    isAcmeAlpnConfigured
    isAcmeDnsConfigured
    isApex
    isConfigured
    isWildcard
  }
}
`
)

// ListAppCertificates returns all the app certificates
func ListAppCertificates(
	ctx context.Context,
	client graphql.Client,
	options *ListAppCertificatesRequestConfiguration,
) (*ListAppCertificatesResponse, error) {

	// Check for options
	variables := &__ListAppCertificatesInput{}
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
		OpName:    "ListAppCertificates",
		Query:     queryAppCertificateList,
		Variables: variables,
	}
	var err error

	var data ListAppCertificatesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// GetAppCertificate returns the specified certificate
func GetAppCertificate(
	ctx context.Context,
	client graphql.Client,
	Id string,
) (*GetAppCertificateResponse, error) {
	req := &graphql.Request{
		OpName: "GetAppCertificate",
		Query:  queryAppCertificateGet,
		Variables: &__GetAppCertificateInput{
			Id: Id,
		},
	}
	var err error

	var data GetAppCertificateResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
