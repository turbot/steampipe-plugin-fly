---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/fly.svg"
brand_color: "#8B5CF6"
display_name: "Fly.io"
short_name: "fly"
description: "Steampipe plugin to query applications, volumes, databases, and more from your Fly organization."
og_description: "Query Fly.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/fly-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Fly.io + Steampipe

[Fly.io](https://fly.io) provides a global application distribution platform where you can run your code in Firecracker microVMs worldwide.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List all apps deployed in your organization:

```sql
select
  name,
  app_url,
  hostname,
  status
from
  fly_app;
```

```
+---------------------------+-----------------------------+-----------------------------------+-----------+
| name                      | app_url                     | hostname                          | status    |
+---------------------------+-----------------------------+-----------------------------------+-----------+
| silent-meadow-6123        | https://2a09:8280:1::1:c64a | silent-meadow-6123.fly.dev        | running   |
| fly-builder-icy-tree-3230 | <null>                      | fly-builder-icy-tree-3230.fly.dev | suspended |
+---------------------------+-----------------------------+-----------------------------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/fly/tables)**

## Get started

### Install

Download and install the latest Fly plugin:

```bash
steampipe plugin install fly
```

### Credentials

| Item        | Description                                                                                                                                                            |
| ----------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Fly requires an [API token](https://fly.io/docs/flyctl/auth-token/) for all requests.                                                                                  |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the API token permissions also change.                          |
| Radius      | Each connection represents a single Fly Installation.                                                                                                                  |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/fly.spc`)<br />2. Credentials specified in environment variable, e.g., `FLY_API_TOKEN`. |

### Configuration

Installing the latest fly plugin will create a config file (`~/.steampipe/config/fly.spc`) with a single connection named `fly`:

```hcl
connection "fly" {
  plugin = "fly"

  # Fly.io API token.
  # To generate the token please visit https://fly.io/docs/flyctl/auth-token/
  # This can also be set via the `FLY_API_TOKEN` environment variable.
  # api_token = "97GtVsdAPwowRToaWDtgZtILdXI_agszONwajQslZ1o"
}
```

### Credentials from Environment Variables

The Fly plugin will use the standard Fly environment variables to obtain credentials **only if other argument (`api_token`) is not specified** in the connection:

```sh
export FLY_API_TOKEN=97GtVsdAPwowRToaWDtgZtILdXI_agszONwajQslZ1o
```


