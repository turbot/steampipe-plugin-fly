---
title: "Steampipe Table: fly_location - Query Fly Locations using SQL"
description: "Allows users to query Fly Locations, specifically the details about regional deployment of applications."
---

# Table: fly_location - Query Fly Locations using SQL

Fly Locations refers to the geographical locations where Fly.io applications are deployed. Fly.io provides a platform for building and running applications close to users, with a global application network that enables developers to run their code anywhere. The locations represent the regions where these applications are running, providing insights into the distribution and reach of the applications. 

## Table Usage Guide

The `fly_location` table provides insights into the geographical deployment of applications on Fly.io. As a DevOps engineer, explore location-specific details through this table, including region codes, names, and associated metadata. Utilize it to understand the geographical distribution of your applications, helping in making informed decisions about scaling and resource allocation.

## Examples

### Basic info
Explore the geographical details of flight locations, such as name, title, and locality, to gain a better understanding of your flight data and make informed decisions about travel routes and destinations.

```sql
select
  name,
  title,
  locality,
  state,
  country
from
  fly_location;
```

### List all available location in a specific country
Discover the segments that are available within a specific country. This can be useful for pinpointing locations for potential business expansion or understanding demographic distribution.

```sql
select
  title,
  locality,
  state,
  coordinates
from
  fly_location
where
  country = 'India';
```

### Get count of Fly data centers per country
Analyze the distribution of Fly data centers across different countries. This is useful for understanding where resources are concentrated and can inform decisions about where to deploy additional infrastructure.

```sql
select
  country,
  count(name)
from
  fly_location
group by
  country;
```