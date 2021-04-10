package beta

import (
	"context"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

type DirectoryRolesService service

type DirectoryRolesResponse struct {
	OData
	DirectoryRoles []*DirectoryRole `json:"value"`
}

// ListAll lists the directory roles that are activated in the tenant.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-list?view=graph-rest-beta&tabs=http
func (s *DirectoryRolesService) ListAll(ctx context.Context) (*DirectoryRolesResponse, error) {
	u := msgraph.URL(DirectoryRoles).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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
	OData
	DirectoryRole
}

// GetDirectoryRole retrieves the properties of a directoryRole object. The role must be activated in tenant for a successful response.
//
// GET /directoryRoles/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-get?view=graph-rest-beta&tabs=http
func (s *DirectoryRolesService) GetDirectoryRole(ctx context.Context, directoryRoleID string) (*DirectoryRoleResponse, error) {
	u := msgraph.URL(DirectoryRoles, directoryRoleID).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

// ActivateDirectoryRole activates a directory role. To read a directory role or update its members, it must first be activated in the tenant.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-post-directoryroles?view=graph-rest-beta&tabs=http
func (s *DirectoryRolesService) ActivateDirectoryRole(ctx context.Context, roleTemplateID string) (*DirectoryRoleResponse, error) {
	directoryRole := new(DirectoryRole)
	directoryRole.RoleTemplateId = msgraph.String(roleTemplateID)

	u := msgraph.URL(DirectoryRoles).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, directoryRole)
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

// AddMember creates a new directory role member.
// You can use both the object ID and template ID of the directoryRole with this API. The template ID of a built-in role
// is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
//
// POST /directoryRoles/{id}/members/$ref
//
// objectOrTemplateIDOfDirectoryRole - the object ID and template ID of the directoryRole
// directoryObjectID - the directory object ID e.g.  user ID, group ID etc.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/directoryrole-post-members?view=graph-rest-beta&tabs=http
func (s *DirectoryRolesService) AddMember(ctx context.Context, objectOrTemplateIDOfDirectoryRole, directoryObjectID string) error {
	requestData := map[string]string{
		"@odata.id": msgraph.URL(DirectoryObjects, directoryObjectID).BuildWithPrefix(defaultBaseURL),
	}

	u := msgraph.URL(DirectoryRoles, objectOrTemplateIDOfDirectoryRole).Append(Members).Append(Refs).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, requestData)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
