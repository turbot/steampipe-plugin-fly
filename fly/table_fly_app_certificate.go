package fly

import (
	"context"

	"github.com/turbot/steampipe-plugin-fly/flyapi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	provider "github.com/fly-apps/terraform-provider-fly/graphql"
)

//// TABLE DEFINITION

func tableFlyAppCertificate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fly_app_certificate",
		Description: "Fly App Certificate",
		List: &plugin.ListConfig{
			ParentHydrate: listFlyApps,
			Hydrate:       listFlyAppCertificates,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFlyAppCertificate,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "domain", Description: "The fully qualified domain name of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "A unique identifier of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "hostname", Description: "The hostname of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "The timestamp when the certificate was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "source", Description: "The source of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "certificate_authority", Description: "The certificate authority.", Type: proto.ColumnType_STRING},
			{Name: "verified", Description: "If true, the certificate DNS configuration is verified.", Type: proto.ColumnType_BOOL},
			{Name: "client_status", Description: "The client status of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "dns_provider", Description: "The DNS provider of the certificate.", Type: proto.ColumnType_STRING},
			{Name: "dns_validation_hostname", Description: "Specifies the DNS validation hostname.", Type: proto.ColumnType_STRING},
			{Name: "dns_validation_instructions", Description: "Specifies the DNS validation instructions.", Type: proto.ColumnType_STRING},
			{Name: "dns_validation_target", Description: "Specifies the DNS validation target.", Type: proto.ColumnType_STRING},
			{Name: "is_acme_alpn_configured", Description: "If true, certificate acme-alpn is configured.", Type: proto.ColumnType_BOOL},
			{Name: "is_acme_dns_configured", Description: "If true, acme-dns is configured for domain validation.", Type: proto.ColumnType_BOOL},
			{Name: "is_apex", Description: "True, if the certificate is a apex certificate.", Type: proto.ColumnType_BOOL},
			{Name: "is_configured", Description: "If true, certificate is configured with valid DNS configuration.", Type: proto.ColumnType_BOOL},
			{Name: "is_wildcard", Description: "If true, the hostname of the certificate contains wildcard.", Type: proto.ColumnType_BOOL},
		},
	}
}

//// LIST FUNCTION

func listFlyAppCertificates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	appData := h.Item.(provider.GetFullAppApp)
	appID := appData.Name

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app_certificate.listFlyAppCertificates", "connection_error", err)
		return nil, err
	}

	options := &flyapi.ListAppCertificatesRequestConfiguration{
		AppId: appID,
	}

	// There is no max page limit as such defined
	// but, we set the default page limit as 5000
	pageLimit := 5000

	// Adjust page limit, if less than default value
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < pageLimit {
			pageLimit = int(*limit)
		}
	}
	options.Limit = pageLimit

	for {
		query, err := flyapi.ListAppCertificates(context.Background(), conn.Graphql, options)
		if err != nil {
			plugin.Logger(ctx).Error("fly_app_certificate.listFlyAppCertificates", "query_error", err)
			return nil, err
		}

		for _, cert := range query.App.Certificates.Nodes {
			d.StreamListItem(ctx, cert)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all resources are processed
		if !query.App.Certificates.PageInfo.HasNextPage {
			break
		}

		// Else set the next page cursor
		options.EndCursor = query.App.Certificates.PageInfo.EndCursor
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getFlyAppCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	certID := d.EqualsQualString("id")
	
	// Return nil, if empty
	if certID == "" {
		return nil, nil
	}

	// Create client
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app_certificate.getFlyAppCertificate", "connection_error", err)
		return nil, err
	}

	query, err := flyapi.GetAppCertificate(context.Background(), conn.Graphql, certID)
	if err != nil {
		plugin.Logger(ctx).Error("fly_app_certificate.getFlyAppCertificate", "query_error", err)
		return nil, err
	}

	return query.Certificate, nil
}
