package {{cookiecutter.plugin_name}}

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func table{{cookiecutter.table_name}}(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "{{cookiecutter.plugin_name}}_{{ cookiecutter.table_name }}",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: {{cookiecutter.table_name}},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "text", Require: plugin.Optional},
				{Name: "settings", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
            // Result columns
			{Name: "results", Type: proto.ColumnType_STRING, Transform: transform.FromField("results"), Description: "results"},

            // Input Columns
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromQual("text"), Description: "text"},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the completion API request parameters."},
		},
	}
}

type {{ cookiecutter.table_name }}RequestQual struct {
	Text *string  `json:"text,omitempty"`
}

type {{ cookiecutter.table_name }}Row struct {
	Result string
}

func {{cookiecutter.table_name}}(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

    // NOTE: IMPLEMENT THIS based on the service/provider
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("{{cookiecutter.plugin_name}}.{{cookiecutter.table_name}}", "connection_error", err)
		return nil, err
	}


	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		// Overwrite any settings provided in the settings qual. If a field
		// is not passed in the settings, then default to the settings above.
		var crQual {{ cookiecutter.table_name }}RequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("{{cookiecutter.plugin_name}}.{{cookiecutter.table_name}}", "unmarshal_error", err)
			return nil, err
		}
	}

    // Query the sdk with appropriate methods and serialize the response
	//resp, err := conn.SOME SDK/Client METHOD(cr)
	if err != nil {
		plugin.Logger(ctx).Error("{{cookiecutter.plugin_name}}.{{cookiecutter.table_name}}", "api_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Trace("{{cookiecutter.plugin_name}}_{{cookiecutter.table_name}}", "response", resp)
	for _, c := range resp.RESULTS {
		row := {{cookiecutter.table_name}}Row{c, cr.Text}
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}
