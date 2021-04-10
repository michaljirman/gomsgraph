// +build integration

package integration

import (
	"context"

	"github.com/michaljirman/gomsgraph/msgraph/beta"
)

var (
	client *beta.Client
)

func init() {
	ctx := context.Background()
	client = beta.NewDefaultClient(ctx)
}
