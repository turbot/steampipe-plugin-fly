![image](https://hub.steampipe.io/images/plugins/turbot/fly-social-graphic.png)

# Fly Plugin for Steampipe

Use SQL to query applications, volumes, databases, and more from your Fly organization.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/fly)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/fly/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-fly/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install fly
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/fly#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/fly#configuration).

Run a query:

```sql
select
  name,
  app_url,
  hostname,
  status
from
  fly_app;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-fly.git
cd steampipe-plugin-fly
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/fly.spc
```

Try it!

```shell
steampipe query
> .inspect fly
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-fly/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Fly Plugin](https://github.com/turbot/steampipe-plugin-fly/labels/help%20wanted)
