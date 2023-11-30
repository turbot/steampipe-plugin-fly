---
title: "Steampipe Table: fly_app - Query OCI Fly Apps using SQL"
description: "Allows users to query Fly Apps, specifically the app details, providing insights into app configurations and statuses."
---

# Table: fly_app - Query OCI Fly Apps using SQL

Fly App is a resource within Oracle Cloud Infrastructure (OCI) that allows you to deploy and manage applications. It provides a platform for developers to build, deploy, and scale applications in a consistent manner. Fly App helps you manage the lifecycle of your applications, ensuring they are always up-to-date and running efficiently.

## Table Usage Guide

The `fly_app` table provides insights into Fly Apps within Oracle Cloud Infrastructure (OCI). As a developer or a DevOps engineer, explore app-specific details through this table, including configurations, statuses, and associated metadata. Utilize it to uncover information about apps, such as their current status, configurations, and other crucial details that can help in managing and scaling applications efficiently.

## Examples

### Basic info
Explore which applications are currently active, by assessing their status and associated URLs. This can help in managing and monitoring the applications more effectively.

```sql
select
  name,
  app_url,
  status,
  hostname
from
  fly_app;
```

### List suspended apps
Explore which applications have been suspended, providing a way to manage and review their status and further action. This is useful for maintaining the health and efficiency of your application environment.

```sql
select
  name,
  app_url,
  status,
  hostname
from
  fly_app
where
  status = 'suspended';
```

### List unencrypted volumes attached to the instances
This query helps to identify any unencrypted volumes attached to applications, which is crucial for maintaining data security and compliance. It's a practical tool for reviewing and rectifying potential vulnerabilities in your system.

```sql
select
  a.name,
  v.name,
  v.encrypted
from
  fly_app as a
  join fly_volume as v on v.app_id = a.id and not v.encrypted;
```

### List apps with unverified certificates
Discover the segments that contain applications with unverified certificates. This is useful in identifying potential security risks within your system.

```sql
select
  a.name as app_name,
  a.status,
  c.domain,
  c.hostname
from
  fly_app as a
  join fly_app_certificate as c on a.id = c.app_id and not c.verified;
```