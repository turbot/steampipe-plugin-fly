---
title: "Steampipe Table: fly_ip - Query Fly IPs using SQL"
description: "Allows users to query Fly IPs, specifically the IP ranges and their associated regions, providing insights into network distribution and potential anomalies."
---

# Table: fly_ip - Query Fly IPs using SQL

Fly IP is a resource within Fly.io that represents the IP ranges allocated to the Fly.io platform. These IP ranges are distributed across various regions, enabling the deployment of applications closer to end users for better performance. Fly IP provides critical information about the geographical distribution of your applications and their network accessibility.

## Table Usage Guide

The `fly_ip` table provides insights into the IP ranges within Fly.io. As a network administrator, explore IP-specific details through this table, including the associated regions, and the network distribution of your applications. Utilize it to uncover information about the geographical spread of your applications, the network accessibility, and to plan for better performance and lower latency.

## Examples

### Basic info
Explore the creation dates and types of IP addresses within different regions. This allows the user to gain a clearer overview of their network infrastructure and aids in managing and optimizing resource allocation.

```sql+postgres
select
  address,
  id,
  created_at,
  type,
  region
from
  fly_ip;
```

```sql+sqlite
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
Discover the segments that have IP addresses allocated to applications that are currently suspended. This is beneficial in identifying potential security risks or freeing up resources.

```sql+postgres
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

```sql+sqlite
select
  ip.address,
  ip.type as ip_type,
  a.name as app_name
from
  fly_app as a,
  json_each(a.ip_addresses, '$.nodes') as node
  join fly_ip as ip on ip.id = json_extract(node.value, '$.id')
where
  a.status = 'suspended';
```

### List IP allocation information
Explore which applications are associated with specific IP addresses and their types. This can be useful in understanding the distribution and utilization of IP addresses across different applications.

```sql+postgres
select
  ip.address,
  ip.type as ip_type,
  a.name as app_name
from
  fly_app as a,
  jsonb_array_elements(ip_addresses -> 'nodes') as node
  join fly_ip as ip on ip.id = node ->> 'id';
```

```sql+sqlite
select
  ip.address,
  ip.type as ip_type,
  a.name as app_name
from
  fly_app as a,
  json_each(a.ip_addresses, '$.nodes') as node
  join fly_ip as ip on ip.id = json_extract(node.value, '$.id');
```