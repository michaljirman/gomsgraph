package core

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/microsoft"
)

const (
	scopesDefault = "https://graph.microsoft.com/.default"
)

func GetOAuth2Client(ctx context.Context, tenantID string, clientID string, clientSecret string) *http.Client {
	microsoftEndpoints := microsoft.AzureADEndpoint(tenantID)
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{scopesDefault},
		TokenURL:     microsoftEndpoints.TokenURL,
	}

	return conf.Client(ctx)
}
