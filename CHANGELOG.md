## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#14](https://github.com/turbot/steampipe-plugin-fly/pull/14))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#12](https://github.com/turbot/steampipe-plugin-fly/pull/12))
- Recompiled plugin with Go version `1.21`. ([#12](https://github.com/turbot/steampipe-plugin-fly/pull/12))

## v0.1.1 [2023-07-11]

_Bug fixes_

- Fixed the plugin's config argument to use `api_token` instead of `fly_api_token` to align with the API documentation. ([#6](https://github.com/turbot/steampipe-plugin-fly/pull/6))

## v0.1.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#4](https://github.com/turbot/steampipe-plugin-fly/pull/4))

## v0.0.1 [2023-01-18]

_What's new?_

- New tables added
  - [fly_app](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_app)
  - [fly_app_certificate](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_app_certificate)
  - [fly_ip](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_ip)
  - [fly_location](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_location)
  - [fly_machine](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_machine)
  - [fly_organization](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_organization)
  - [fly_organization_member](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_organization_member)
  - [fly_redis_database](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_redis_database)
  - [fly_volume](https://hub.steampipe.io/plugins/turbot/fly/tables/fly_volume)
