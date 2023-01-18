# Table: fly_organization

An organization is a way of sharing applications and resources between Fly.io users.

## Examples

### Basic info

```sql
select
  name,
  id,
  slug,
  type,
  billing_status
from
  fly_organization;
```

### List organizations with no payment method configured

```sql
select
  name,
  id,
  slug,
  type,
  billing_status
from
  fly_organization
where
  not is_credit_card_saved;
```

### List organizations without SSH certificate

```sql
select
  name,
  id,
  slug,
  type
from
  fly_organization
where
  ssh_certificate is null;
```
