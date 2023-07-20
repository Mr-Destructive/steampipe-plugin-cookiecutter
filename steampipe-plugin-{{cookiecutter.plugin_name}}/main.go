package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "github.com/turbot/steampipe-plugin-{{cookiecutter.plugin_name}}/{{ cookiecutter.plugin_name }}"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: {{cookiecutter.plugin_name}}.Plugin})
}
