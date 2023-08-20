
# {{cookiecutter.plugin_standard_name}} Plugin for Steampipe

Use SQL to query {{cookiecutter.table_name}} and more from {{cookiecutter.plugin_standard_name}}.

- **[Get started →](https://hub.steampipe.io/plugins/{{cookiecutter.github_username}}/{{cookiecutter.plugin_name}})**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/{{cookiecutter.github_username}}/{{cookiecutter.plugin_name}}/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/{{cookiecutter.github_username}}/steampipe-plugin-{{cookiecutter.plugin_name}}/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
{%- if cookiecutter.github_username == "turbot" -%}
steampipe plugin install {{cookiecutter.plugin_name}}
{%- else -%}
steampipe plugin install {{cookiecutter.github_username}}/{{cookiecutter.plugin_name}}
{%- endif -%}
```

Configure your plugin in `~/.steampipe/config/{{cookiecutter.plugin_name}}.spc`:

```hcl
connection "{{cookiecutter.plugin_name}}" {
  plugin  = "{{cookiecutter.github_username}}/{{cookiecutter.plugin_name}}"
  # ...
}
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  *
from
  {{cookiecutter.table_name}}
where
  something = ''
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/{{cookiecutter.github_username}}/steampipe-plugin-{{cookiecutter.plugin_name}}.git
cd steampipe-plugin-{{cookiecutter.plugin_name}}
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/{{cookiecutter.plugin_name}}.spc
```

Try it!

```
steampipe query
> .inspect {{cookiecutter.plugin_name}}
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-{{cookiecutter.plugin_name}}/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [{{cookiecutter.plugin_standard_name}}Plugin](https://github.com/{{cookiecutter.github_username}}/steampipe-plugin-{{cookiecutter.plugin_name}}/labels/help%20wanted)
