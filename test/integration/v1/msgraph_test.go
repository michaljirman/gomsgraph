// +build integration

package v1

import (
	"context"
	"log"
	"os"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	msgraph "github.com/michaljirman/gomsgraph/msgraph/v1"
)

var (
	client *msgraph.Client
)

func init() {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("env variable CLIENT_ID is required for integration tests")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("env variable CLIENT_SECRET is required for integration tests")
	}

	clientTenantID := os.Getenv("CLIENT_TENANT_ID")
	if clientTenantID == "" {
		log.Fatal("env variable CLIENT_TENANT_ID is required for integration tests")
	}

	ctx := context.Background()
	tc := core.GetOAuth2Client(ctx, clientTenantID, clientID, clientSecret)
	client = msgraph.NewClient(tc)
}
