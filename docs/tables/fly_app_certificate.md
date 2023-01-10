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
