# Table: fly_location

Fly Location is a collection of regions where Fly has edges and/or data centers.

## Examples

### Basic info

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

### Get count of Fly data centers per country

```sql
select
  country,
  count(name)
from
  fly_location
group by
  country;
```
