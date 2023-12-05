---
title: "Steampipe Table: fly_app_certificate - Query Fly.io App Certificates using SQL"
description: "Allows users to query App Certificates in Fly.io, specifically the certificate details for each application, providing insights into application security and SSL/TLS configurations."
---

# Table: fly_app_certificate - Query Fly.io App Certificates using SQL

Fly.io App Certificates are part of the Fly.io platform that allows you to handle SSL/TLS for your applications. These certificates are automatically managed and renewed, ensuring your applications are always served over HTTPS. They play a crucial role in maintaining the security and integrity of data transmitted between your Fly.io applications and their users.

## Table Usage Guide

The `fly_app_certificate` table provides insights into App Certificates within Fly.io. As a DevOps engineer, explore certificate-specific details through this table, including expiration dates, issuing authorities, and associated metadata. Utilize it to monitor your SSL/TLS configurations, ensure certificates are up to date, and maintain the security of your Fly.io applications.

## Examples

### Basic info
Explore which domain certificates were created and their respective sources. This can help identify the origin of each certificate, providing insights into potential security risks or issues.

```sql+postgres
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate;
```

```sql+sqlite
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate;
```

### List unverified certificates
Discover the segments that contain unverified certificates within your application. This could be useful, for example, in identifying potential security risks and ensuring that all certificates are valid and up-to-date.

```sql+postgres
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where
  not verified;
```

```sql+sqlite
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where
  not verified;
```

### List certificates that do not have valid DNS configuration
Identify instances where certain certificates may have been created without a valid DNS configuration. This can be useful for troubleshooting connectivity issues or ensuring proper setup of digital certificates.

```sql+postgres
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where
  not is_configured;
```

```sql+sqlite
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where
  is_configured = 0;
```

### List DNS configuration details of certificates
Explore which domain certificates have certain DNS configurations. This can be particularly useful for understanding how your certificates are set up and where potential configuration issues may lie.

```sql+postgres
select
  domain,
  id,
  dns_provider,
  dns_validation_hostname,
  dns_validation_instructions,
  dns_validation_target
from
  fly_app_certificate;
```

```sql+sqlite
select
  domain,
  id,
  dns_provider,
  dns_validation_hostname,
  dns_validation_instructions,
  dns_validation_target
from
  fly_app_certificate;
```

### List certificates associated with a specific app
Explore which certificates are linked to a specific application, allowing you to assess the security elements within your application's configuration. This could be beneficial in identifying areas where updates or changes may be necessary for compliance or improved security.

```sql+postgres
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where 
  app_id = 'fly-builder-purple-cloud-1058';
```

```sql+sqlite
select
  domain,
  id,
  hostname,
  created_at,
  source
from
  fly_app_certificate
where 
  app_id = 'fly-builder-purple-cloud-1058';
```