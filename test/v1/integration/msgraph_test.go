// +build integration

package integration

import (
	"context"

	v1 "github.com/michaljirman/gomsgraph/msgraph/v1"
)

var (
	client *v1.Client
)

func init() {
	ctx := context.Background()
	client = v1.NewDefaultClient(ctx)
}
