---
title: "Steampipe Table: fly_organization_member - Query Fly.io Organization Members using SQL"
description: "Allows users to query Organization Members in Fly.io, providing details on member roles, organizations they belong to, and their associated metadata."
---

# Table: fly_organization_member - Query Fly.io Organization Members using SQL

Fly.io is a deployment platform that automates the work of managing and orchestrating containers. It provides a global application runtime for developers to deploy and run application code in any language, using datacenters closer to users for better performance. An Organization Member in Fly.io is an individual who has been granted access to an organization and may have different roles and permissions based on their assignment.

## Table Usage Guide

The `fly_organization_member` table provides insights into Organization Members within Fly.io. As a DevOps engineer or IT administrator, you can use this table to understand the roles and permissions of each member within your organization. You can also use it to monitor and manage members' access to different resources, ensuring the security and efficiency of your organization's operations.

## Examples

### Basic info
Explore which roles each member holds within your organization, gaining insights into the distribution of responsibilities and hierarchical structure. This can be particularly useful for assessing the elements within your team and identifying areas for growth or reorganization.

```sql+postgres
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member;
```

```sql+sqlite
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
Explore which members in your organization hold administrative roles. This helps in understanding the distribution of administrative privileges and aids in managing user permissions effectively.

```sql+postgres
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

```sql+sqlite
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

### List all members with two-factor authentication disabled
Discover the members who have not enabled two-factor authentication. This is useful to identify potential security risks and ensure that all members are adhering to best practices for account protection.

```sql+postgres
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

```sql+sqlite
select
  name,
  email,
  username,
  role,
  organization_id
from
  fly_organization_member
where
  two_factor_protection = 0;
```

### List all members with restricted access
Identify members within an organization who have restricted access, providing insights into user roles and permissions for better access management.

```sql+postgres
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

```sql+sqlite
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