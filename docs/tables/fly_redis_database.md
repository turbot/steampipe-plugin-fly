# Table: fly_redis_database

Redis by Upstash is a fully-managed, Redis-compatible database service offering global read replicas for reduced latency in distant regions. Upstash databases are provisioned inside your Fly organization, ensuring private, low-latency connections to your Fly applications.

## Examples

### Basic info

```sql
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database;
```

### List databases with no replica

```sql
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  read_regions is null
  or jsonb_array_length(read_regions) = 0;
```

### List databases with object eviction enabled

```sql
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  (options -> 'eviction')::boolean;
```

### List large databases

```sql
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  addOnPlan ->> 'maxDataSize' = '3 GB';
```
