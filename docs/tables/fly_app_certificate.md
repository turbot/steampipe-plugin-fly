# Table: fly_app_certificate

The `fly_app_certificate` table can be used to query information about the certificates associated with a deployed application.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

### List DNS configuration details of certificates

```sql
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
