# Table: fly_ip

The `fly_ip` table can be used to query information about IP addresses allocated to the application.

## Examples

### Basic info

```sql
select
  address,
  id,
  created_at,
  type,
  region
from
  fly_ip;
```

### List IP addresses allocated to suspended apps

```sql
select
  ip.address,
  ip.type as ip_type,
  a.name as app_name
from
  fly_app as a,
  jsonb_array_elements(ip_addresses -> 'nodes') as node
  join fly_ip as ip on ip.id = node ->> 'id'
where
  a.status = 'suspended';
```

### List IP allocation information

```sql
select
  ip.address,
  ip.type as ip_type,
  a.name as app_name
from
  fly_app as a,
  jsonb_array_elements(ip_addresses -> 'nodes') as node
  join fly_ip as ip on ip.id = node ->> 'id';
```
