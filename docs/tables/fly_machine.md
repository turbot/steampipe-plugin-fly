---
title: "Steampipe Table: fly_machine - Query Fly.io Machines using SQL"
description: "Allows users to query Machines on Fly.io, providing details about the virtual machines including their status, region, and associated applications."
---

# Table: fly_machine - Query Fly.io Machines using SQL

Fly.io is a platform for running application servers close to users. It provides a global application runtime that allows deployment of applications to any city in the world. This service enables users to run their applications on a collection of virtual machines called 'Machines'.

## Table Usage Guide

The `fly_machine` table provides insights into Machines within the Fly.io platform. As a developer or system administrator, explore machine-specific details through this table, including their status, region, and associated applications. Utilize it to uncover information about machines, such as those in a particular region, the status of machines, and the applications they are associated with.

## Examples

### Basic info
Explore the status and location of various machines in your fleet to understand their operational distribution and longevity. This could be useful for assessing the need for new purchases or redistributions based on regional demands and machine age.

```sql
select
  name,
  id,
  state,
  region,
  created_at
from
  fly_machine;
```

### List stopped machines
Discover the segments that consist of halted machines, allowing you to analyze your resources and optimize accordingly. This is useful in managing resource allocation and preventing unnecessary costs associated with idle machines.

```sql
select
  name,
  id,
  state,
  region,
  created_at
from
  fly_machine
where
  state = 'stopped';
```

### List machines by app
Discover the segments that are associated with a specific application by analyzing the state, region, and creation date of each machine. This can be useful in managing resources and tracking the performance of different applications.

```sql
select
  name,
  id,
  state,
  region,
  created_at
from
  fly_machine
where
  app_id = 'fly-builder-icy-tree-3230';
```

### List unencrypted volumes attached to the machines
Discover the segments that consist of machines with unencrypted volumes attached to them. This is beneficial for identifying potential security risks in your system.

```sql
select
  m.name as machine,
  v.name as volume,
  v.encrypted
from
  fly_machine as m,
  jsonb_array_elements(config -> 'mounts') as mount
  join fly_volume as v on v.id = mount ->> 'volume' and not v.encrypted;
```