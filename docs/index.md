---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/fly.svg"
brand_color: "#24185B"
display_name: "Fly.io"
short_name: "fly"
description: "Steampipe plugin tto query applications, volumes, and more from your Fly organization."
og_description: "Query Fly.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/fly-social-graphic.png"
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

The API token can be generated using the [CLI](https://fly.io/docs/flyctl/auth-token) to the Fly.io; or from the [organization dashboard](https://fly.io/dashboard) itself.

To generate the API token from the [organization dashboard](https://fly.io/dashboard):

- Log into the [Fly.io](https://fly.io).
- Go to the **Account** section in the top-right of the Console, and select **Access Tokens** from the drop-down.
- Create token by providing a name.

### Configuration

Installing the latest fly plugin will create a config file (`~/.steampipe/config/fly.spc`) with a single connection named `fly`:

```hcl
connection "fly" {
  plugin = "fly"

  # The fly.io api token.
  token = "YOUR_ACCESS_TOKEN"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-fly
- Community: [Slack Channel](https://steampipe.io/community/join)
