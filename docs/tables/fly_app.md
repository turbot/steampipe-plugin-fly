# Table: fly_app

An app groups one or more VM instances into a single administrative entity. Fly.io allows you to deploy any kind of app as long as it is packaged in a Docker image.

## Examples

### Basic info

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

```sql
select
  a.name,
  v.name,
  v.encrypted
from
  fly_app as a
  join fly_volume as v on v.app_id = a.id and not v.encrypted;
```
