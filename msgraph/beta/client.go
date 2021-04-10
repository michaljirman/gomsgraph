package beta

import (
	"context"
	"net/http"
	"net/url"

	"github.com/michaljirman/gomsgraph/msgraph"
)

const (
	defaultBaseURL = "https://graph.microsoft.com/beta/"
	userAgent      = "gomsgraph"
)

// A Client manages communication with the MS Graph API.
type Client struct {
	msgraph.BaseClient

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	Users                  *UsersService
	Groups                 *GroupsService
	DirectoryRoles         *DirectoryRolesService
	DirectoryRoleTemplates *DirectoryRoleTemplatesService
	SchemaExtensions       *SchemaExtensionsService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{}
	c.BaseClient.Client = httpClient
	c.BaseClient.BaseURL = baseURL
	c.BaseClient.UserAgent = userAgent
	c.common.client = c
	c.Users = (*UsersService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)
	c.DirectoryRoles = (*DirectoryRolesService)(&c.common)
	c.DirectoryRoleTemplates = (*DirectoryRoleTemplatesService)(&c.common)
	c.SchemaExtensions = (*SchemaExtensionsService)(&c.common)
	return c
}

// NewDefaultClient creates a new client with default configuration for MS Graph Beta API
func NewDefaultClient(ctx context.Context) *Client {
	return NewClient(msgraph.NewOAuth2ClientFromEnvsOrFail(ctx))
}
