# Table: fly_organization_member

The `fly_organization_member` table can be used to query information about members of an organization.

## Examples

### Basic info

```sql
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member;
```

### List all admins in the organization

```sql
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member
where
  role = 'ADMIN';
```

### List all members for which two-factor authentication is not configured

```sql
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member
where
  not two_factor_protection;
```

### List all members with restricted access

```sql
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member
where
  trust = 'RESTRICTED';
```
