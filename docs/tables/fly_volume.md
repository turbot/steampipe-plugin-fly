---
title: "Steampipe Table: fly_volume - Query Fly.io Volumes using SQL"
description: "Allows users to query Fly.io Volumes, providing detailed information about the volumes attached to apps."
---

# Table: fly_volume - Query Fly.io Volumes using SQL

Fly.io Volumes are block storage devices that can be attached to apps for persistent data storage. They provide a way to store data that persists across deployments and restarts, and can be shared between instances of an app. Volumes are attached to a specific region and can be moved between regions if needed.

## Table Usage Guide

The `fly_volume` table provides insights into volumes within Fly.io. As a developer or system administrator, explore volume-specific details through this table, including the attached app, size, and region. Utilize it to manage and monitor your persistent data storage across different regions and applications.

## Examples

### Basic info
Explore which volumes in your system are encrypted and their respective statuses across different regions. This can help secure your data by identifying any unencrypted volumes that may pose a security risk.

```sql+postgres
select
  name,
  id,
  status,
  encrypted,
  region
from
  fly_volume;
```

```sql+sqlite
select
  name,
  id,
  status,
  encrypted,
  region
from
  fly_volume;
```

### List unencrypted volumes
Uncover the details of unencrypted volumes within your system to enhance your data security. This query provides a tool to identify potential vulnerabilities and take corrective actions.

```sql+postgres
select
  name,
  id,
  status,
  region
from
  fly_volume
where
  not encrypted;
```

```sql+sqlite
select
  name,
  id,
  status,
  region
from
  fly_volume
where
  encrypted = 0;
```

### List unused volumes
Identify unutilized storage volumes that are not attached to any machine, providing insights into potential cost savings or optimization opportunities in your cloud infrastructure.

```sql+postgres
select
  name,
  id,
  status,
  region,
  attached_machine_id
from
  fly_volume
where
  attached_machine_id = '';
```

```sql+sqlite
select
  name,
  id,
  status,
  region,
  attached_machine_id
from
  fly_volume
where
  attached_machine_id = '';
```

### List of volumes with size more than 100GiB
Explore which volumes exceed 100GiB in size. This is beneficial for managing storage resources and identifying potential areas for data reduction or redistribution.

```sql+postgres
select
  name,
  id,
  status,
  region,
  size_gb
from
  fly_volume
where
  size_gb > '100';
```

```sql+sqlite
select
  name,
  id,
  status,
  region,
  size_gb
from
  fly_volume
where
  size_gb > 100;
```

### List volumes attached to suspended applications
Discover the segments that have volumes connected to applications that are currently inactive. This can be useful in identifying resources that may be unnecessarily occupied, thus optimizing resource utilization.

```sql+postgres
select
  v.name,
  v.size_gb,
  a.id
from
  fly_volume as v
  join fly_app as a on v.app_id = a.id and a.status = 'suspended';
```

```sql+sqlite
select
  v.name,
  v.size_gb,
  a.id
from
  fly_volume as v
  join fly_app as a on v.app_id = a.id
where
  a.status = 'suspended';
```