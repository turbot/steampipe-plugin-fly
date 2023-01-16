# Table: fly_volume

A volume on Fly is a slice of a Non-Volatile Memory Express (NVMe) drive on the physical server your app runs on.

## Examples

### Basic info

```sql
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

```sql
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

### List unused volumes

```sql
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

```sql
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

### List volumes attached to suspended applications

```sql
select
  v.name,
  v.size_gb,
  a.id
from
  fly_volume as v
  join fly_app as a on v.app_id = a.id and a.status = 'suspended';
```
