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

### List all available location in a specific country

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

```sql
select
  country,
  count(name)
from
  fly_location
group by
  country;
```
