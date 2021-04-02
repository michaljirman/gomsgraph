package beta

import (
	"context"
	"fmt"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

type DirectoryRoleTemplatesService service

type DirectoryRoleTemplatesResponse struct {
	models.OData
	DirectoryRoleTemplates []*models.DirectoryRoleTemplate `json:"value"`
}

// Retrieve a list of directoryRoleTemplate objects.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryroletemplate-list?view=graph-rest-1.0&tabs=http
func (s *DirectoryRoleTemplatesService) ListAll(ctx context.Context) (*DirectoryRoleTemplatesResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, "directoryRoleTemplates", nil)
	if err != nil {
		return nil, err
	}

	var directoryRoleTemplatesResponse DirectoryRoleTemplatesResponse
	_, err = s.client.Do(ctx, req, &directoryRoleTemplatesResponse)
	if err != nil {
		return nil, err
	}

	return &directoryRoleTemplatesResponse, nil
}

type DirectoryRoleTemplateResponse struct {
	models.OData
	models.DirectoryRoleTemplate
}

// Retrieve the properties and relationships of a directory role template object.
//
// You can use both the object ID and template ID of the directoryRole with this API.
// The template ID of a built-in role is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
//
// GET /directoryRoleTemplates/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-get?view=graph-rest-beta&tabs=http
func (s *DirectoryRoleTemplatesService) GetDirectoryRoleTemplate(ctx context.Context, objectOrTemplateID string) (*DirectoryRoleResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("directoryRoleTemplates/%v", objectOrTemplateID), nil)
	if err != nil {
		return nil, err
	}

	directoryRoleResp := new(DirectoryRoleResponse)
	_, err = s.client.Do(ctx, req, directoryRoleResp)
	if err != nil {
		return nil, err
	}

	return directoryRoleResp, nil
}
