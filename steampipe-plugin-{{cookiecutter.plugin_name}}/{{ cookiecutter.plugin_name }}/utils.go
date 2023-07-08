package {{cookiecutter.plugin_name}}

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// the client is based on the provider/service to be used
//NOTE: It is not necessaty the provider will have a Client object, it could be 
// a different object or logic
func connect(ctx context.Context, d *plugin.QueryData) (*provider.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*provider.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
    // Implement based on the provider
    return nil, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code: 404")
}
