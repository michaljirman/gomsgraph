package beta

import (
	"context"
	"fmt"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph/beta/models"
	"github.com/michaljirman/gomsgraph/msgraph/core"
)

type DirectoryRolesService service

type DirectoryRolesResponse struct {
	models.OData
	DirectoryRoles []*models.DirectoryRole `json:"value"`
}

// List the directory roles that are activated in the tenant.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-list?view=graph-rest-1.0&tabs=http
func (s *DirectoryRolesService) ListAll(ctx context.Context) (*DirectoryRolesResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, "directoryRoles", nil)
	if err != nil {
		return nil, err
	}

	var directoryRolesResponse DirectoryRolesResponse
	_, err = s.client.Do(ctx, req, &directoryRolesResponse)
	if err != nil {
		return nil, err
	}

	return &directoryRolesResponse, nil
}

type DirectoryRoleResponse struct {
	models.OData
	models.DirectoryRole
}

// Retrieve the properties of a directoryRole object. The role must be activated in tenant for a successful response.
//
// GET /directoryRoles/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-get?view=graph-rest-1.0&tabs=http
func (s *DirectoryRolesService) GetDirectoryRole(ctx context.Context, id string) (*DirectoryRoleResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("directoryRoles/%v", id), nil)
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

// Activate a directory role. To read a directory role or update its members, it must first be activated in the tenant.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-post-directoryroles?view=graph-rest-1.0&tabs=http
func (s *DirectoryRolesService) ActivateDirectoryRole(ctx context.Context, roleTemplateID string) (*DirectoryRoleResponse, error) {
	directoryRole := new(models.DirectoryRole)
	directoryRole.RoleTemplateId = core.String(roleTemplateID)

	req, err := s.client.NewRequest(http.MethodPost, "directoryRoles", directoryRole)
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

// Use this API to create a new directory role member.
// You can use both the object ID and template ID of the directoryRole with this API. The template ID of a built-in role
// is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
//
// POST /directoryRoles/{id}/members/$ref
//
// objectOrTemplateIDOfDirectoryRole - the object ID and template ID of the directoryRole
// directoryObjectID - the directory object ID e.g.  user ID, group ID etc.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-post-members?view=graph-rest-1.0&tabs=http
func (s *DirectoryRolesService) AddMember(ctx context.Context, objectOrTemplateIDOfDirectoryRole, directoryObjectID string) error {
	requestData := map[string]string{
		"@odata.id": fmt.Sprintf("%vdirectoryObjects/%s", defaultBaseURL, directoryObjectID),
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("directoryRoles/%s/members/$ref", objectOrTemplateIDOfDirectoryRole), requestData)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
