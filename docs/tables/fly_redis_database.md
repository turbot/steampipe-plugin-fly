---
title: "Steampipe Table: fly_redis_database - Query Fly Redis Databases using SQL"
description: "Allows users to query Fly Redis Databases, specifically the configuration details, providing insights into the database information and potential anomalies."
---

# Table: fly_redis_database - Query Fly Redis Databases using SQL

Fly Redis Database is a service within Fly that allows you to manage and interact with Redis databases. It provides a centralized way to set up and manage databases for various applications, including web applications, microservices, and more. Fly Redis Database helps you stay informed about the health and performance of your databases and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `fly_redis_database` table provides insights into Redis databases within Fly. As a Database Administrator, explore database-specific details through this table, including configuration, status, and associated metadata. Utilize it to uncover information about databases, such as those with specific configurations, the status of databases, and the verification of configuration details.

## Examples

### Basic info
Explore which Redis databases are currently in use, along with their associated hostnames and public URLs. This can be useful in understanding the layout of your resources and identifying the primary region for each database.

```sql+postgres
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database;
```

```sql+sqlite
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database;
```

### List databases with no replica
Discover the segments that consist of databases without replicas. This is useful for identifying potential single points of failure in your system and ensuring data redundancy.

```sql+postgres
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

```sql+sqlite
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  read_regions is null
  or json_array_length(read_regions) = 0;
```

### List databases with object eviction enabled
Explore which Fly Redis databases have the object eviction feature enabled. This query is useful for identifying databases that may require additional memory management due to potential data loss.

```sql+postgres
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

```sql+sqlite
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  json_extract(options, '$.eviction') = 1;
```

### List large databases
Explore which databases are large, specifically those with a maximum data size of 3 GB. This can be useful for identifying databases that may require more storage or management due to their size.

```sql+postgres
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  add_on_plan ->> 'maxDataSize' = '3 GB';
```

```sql+sqlite
select
  name,
  hostname,
  public_url,
  primary_region
from
  fly_redis_database
where
  json_extract(add_on_plan, '$.maxDataSize') = '3 GB';
```