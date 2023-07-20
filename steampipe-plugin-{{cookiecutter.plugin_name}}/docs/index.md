---
organization: Turbot
category: ["ai"]
icon_url: "/images/plugins/turbot/{{cookiecutter.plugin_name}}.svg"
brand_color: "#000000"
display_name: "{{cookiecutter.plugin_standard_name}}"
short_name: "{{cookiecutter.plugin_name}}"
description: "{{cookiecutter.plugin_description}}."
og_description: "Query {{cookiecutter.plugin_standard_name}} with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/{{cookiecutter.plugin_name}}-social-graphic.png"
---

# {{cookiecutter.plugin_standard_name}} + Steampipe


[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Get generations for a given text prompt in your {{cookiecutter.plugin_standard_name}} account:

```sql
select
  
from
  
where
  
```

```
```

## Documentation

- **[Table definitions & examples →](/docs/tables)**

## Get started

### Install

Download and install the latest {{cookiecutter.plugin_standard_name}} plugin:

```bash
steampipe plugin install {{cookiecutter.plugin_name}}
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials |                                                                                                                                                                                  |
| Permissions | API Keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single {{cookiecutter.plugin_name}} Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/{{cookiecutter.plugin_name}}.spc`)<br />2. Credentials specified in environment variables. |

### Configuration

Installing the latest comereai plugin will create a config file (`~/.steampipe/config/{{cookiecutter.permissions}}.spc`) with a single connection named `{{cookiecutter.plugin_name}}`:

```hcl
connection "{{cookiecutter.plugin_name}}" {
  plugin = "{{cookiecutter.plugin_name}}"

}
```


## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-{{cookiecutter.plugin_name}}
- Community: [Slack Channel](https://steampipe.io/community/join)
