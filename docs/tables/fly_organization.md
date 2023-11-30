---
title: "Steampipe Table: fly_organization - Query Fly.io Organizations using SQL"
description: "Allows users to query Fly.io Organizations, providing insights into organization details and associated metadata."
---

# Table: fly_organization - Query Fly.io Organizations using SQL

Fly.io Organizations are a set of users and applications that can be managed collectively on the Fly.io platform. They provide a centralized way to manage and control access to applications and resources. Fly.io Organizations help in efficient collaboration, resource sharing, and access control across different users and applications.

## Table Usage Guide

The `fly_organization` table provides insights into organizations within Fly.io. As a DevOps engineer, explore organization-specific details through this table, including members, applications, and associated metadata. Utilize it to uncover information about organization structure, member roles, and the distribution of applications across different organizations.

## Examples

### Basic info
Explore the basic details of your organization such as its name, unique identifier, type, and billing status. This information can be useful for administrative tasks or tracking billing activities.

```sql
select
  name,
  id,
  slug,
  type,
  billing_status
from
  fly_organization;
```

### List organizations with no payment method configured
Uncover the details of organizations that have not configured a payment method. This is useful for financial auditing or ensuring all organizations in your network have a valid payment method in place.

```sql
select
  name,
  id,
  slug,
  type,
  billing_status
from
  fly_organization
where
  not is_credit_card_saved;
```

### List organizations without SSH certificate
Uncover the details of organizations that lack an SSH certificate, which may indicate a potential security vulnerability. This query is useful for identifying and addressing potential weak points in your organization's security infrastructure.

```sql
select
  name,
  id,
  slug,
  type
from
  fly_organization
where
  ssh_certificate is null;
```