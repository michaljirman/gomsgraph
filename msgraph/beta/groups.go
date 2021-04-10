package beta

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

type GroupsService service

type GroupsResponse struct {
	OData
	Groups []*Group `json:"value"`
}

// ListAll lists all the groups in an organization, including but not limited to Microsoft 365 groups.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-list?view=graph-rest-beta&tabs=http
func (s *GroupsService) ListAll(ctx context.Context, opts *ListOptions) (*GroupsResponse, error) {
	u := msgraph.URL(Groups).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var groupsResponse GroupsResponse
	_, err = s.client.Do(ctx, req, &groupsResponse)
	if err != nil {
		return nil, err
	}

	return &groupsResponse, nil
}

type GroupResponse struct {
	OData
	Group
}

// GetGroup retrieves the properties and relationships of a group object.
//
// GET /groups/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-get?view=graph-rest-beta&tabs=http
func (s *GroupsService) GetGroup(ctx context.Context, groupID string, opts *ListOptions) (*GroupResponse, error) {
	u := msgraph.URL(Groups, groupID).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	groupResp := new(GroupResponse)
	_, err = s.client.Do(ctx, req, groupResp)
	if err != nil {
		return nil, err
	}

	return groupResp, nil
}

// CreateGroup creates a new group as specified in the request body. You can create the following types of groups:
//
// Microsoft 365 group (unified group)
// Security group
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-post-groups?view=graph-rest-beta&tabs=http
func (s *GroupsService) CreateGroup(ctx context.Context, r *Group) (*GroupResponse, error) {
	u := msgraph.URL(Groups).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}

	groupResp := new(GroupResponse)
	_, err = s.client.Do(ctx, req, groupResp)
	if err != nil {
		return nil, err
	}

	return groupResp, nil
}

// UpdateGroup updates the properties of a group object.
//
// PATCH /groups/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-update?view=graph-rest-beta&tabs=http
func (s *GroupsService) UpdateGroup(ctx context.Context, r *Group) error {
	if r.Id == nil {
		return errors.New("group id is required")
	}

	u := msgraph.URL(Groups, *r.Id).Build()
	req, err := s.client.NewRequest(http.MethodPatch, u, r)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup deletes group.
//
// When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days.
// After that time, they are permanently deleted. To learn more, see deletedItems. This applies only to Microsoft 365 groups.
// DELETE /groups/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-delete?view=graph-rest-beta&tabs=http
func (s *GroupsService) DeleteGroup(ctx context.Context, groupID string) error {
	u := msgraph.URL(Groups, groupID).Build()
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListMembers retrieves a list of the group's direct members. A group can have users, organizational contacts, devices,
// service principals and other groups as members. Currently service principals are not listed as group members
// due to staged roll-out of service principals on Graph V1.0 endpoint. This operation is not transitive.
//
// GET /groups/{id}/members
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-list-members?view=graph-rest-beta&tabs=http
func (s *GroupsService) ListMembers(ctx context.Context, groupID string, opts *ListOptions) (*UsersResponse, error) {
	u := msgraph.URL(Groups, groupID).Append(Members).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	usersResp := new(UsersResponse)
	_, err = s.client.Do(ctx, req, usersResp)
	if err != nil {
		return nil, err
	}

	return usersResp, nil
}

// ListOwners retrieves a list of the group's owners. The owners are a set of users who are allowed to
// modify the group object. Owners are currently not available in Microsoft Graph for groups that were created in
// Exchange or groups that are synchronized from an on-premises environment.
//
// GET /groups/{id}/owners
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-list-owners?view=graph-rest-beta&tabs=http
func (s *GroupsService) ListOwners(ctx context.Context, groupID string, opts *ListOptions) (*UsersResponse, error) {
	u := msgraph.URL(Groups, groupID).Append(Owners).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	usersResp := new(UsersResponse)
	_, err = s.client.Do(ctx, req, usersResp)
	if err != nil {
		return nil, err
	}

	return usersResp, nil
}

// AddMember adds a member to a Microsoft 365 group or a security group through the members navigation property.
// You can add users, service principals or other groups.
//
// POST /groups/{group-id}/members/$ref
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-post-members?view=graph-rest-beta&tabs=http
func (s *GroupsService) AddMember(ctx context.Context, groupID, userID string) error {
	requestData := map[string]string{
		"@odata.id": fmt.Sprintf("%vdirectoryObjects/%v", defaultBaseURL, userID),
	}

	u := msgraph.URL(Groups, groupID).Append(Members).Append(Refs).Build()
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

// AddOwner adds a user or service principal to the group's owners. The owners are a set of users or service principals who are allowed to modify the group object.
//
// POST /groups/{id}/owners/$ref
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-post-owners?view=graph-rest-beta&tabs=http
func (s *GroupsService) AddOwner(ctx context.Context, groupID, userID string) error {
	requestData := map[string]string{
		"@odata.id": msgraph.URL(Users, userID).BuildWithPrefix(defaultBaseURL),
	}

	u := msgraph.URL(Groups, groupID).Append(Owners).Append(Refs).Build()
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

// RemoveMember removes a member from a group via the members navigation property.
//
// DELETE /groups/{id}/members/{id}/$ref
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-delete-members?view=graph-rest-beta&tabs=http
func (s *GroupsService) RemoveMember(ctx context.Context, groupID, userID string) error {
	u := msgraph.URL(Groups, groupID).AppendWithID(Members, userID).Append(Refs).Build()
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// RemoveOwner removes an owner from a Microsoft 365 group, a security group, or a mail-enabled security group
// through the owners navigation property. Once owners are assigned to a group, the last owner of the group cannot be removed.
//
// DELETE /groups/{id}/owners/{id}/$ref
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-delete-owners?view=graph-rest-beta&tabs=http
func (s *GroupsService) RemoveOwner(ctx context.Context, groupID, userID string) error {
	u := msgraph.URL(Groups, groupID).AppendWithID(Owners, userID).Append(Refs).Build()
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListAppRoleAssignments retrieves the list of appRoleAssignment that have been granted to a group.
//
// GET /groups/{id}/appRoleAssignments
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-list-approleassignments?view=graph-rest-beta&tabs=http
func (s *GroupsService) ListAppRoleAssignments(ctx context.Context, groupID string, opts *ListOptions) (*AppRoleAssignmentsResponse, error) {
	u := msgraph.URL(Groups, groupID).Append(AppRoleAssignments).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	appRoleAssignmentsResp := new(AppRoleAssignmentsResponse)
	_, err = s.client.Do(ctx, req, appRoleAssignmentsResp)
	if err != nil {
		return nil, err
	}

	return appRoleAssignmentsResp, nil
}

// AddAppRoleAssignments assigns an app role to a group. All direct members of the group will be considered assigned.
// To grant an app role assignment to a group, you need three identifiers:
//
// principalId: The id of the group to which you are assigning the app role.
// resourceId: The id of the resource servicePrincipal which has defined the app role.
// appRoleId: The id of the appRole (defined on the resource service principal) to assign to the group.
//
// POST /groups/{id}/appRoleAssignments
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-post-approleassignments?view=graph-rest-beta&tabs=http
func (s *GroupsService) AddAppRoleAssignments(ctx context.Context, groupID, principalID, resourceID, appRoleID string) (*AppRoleAssignmentResponse, error) {
	requestData := map[string]string{
		"principalId": principalID,
		"resourceId":  resourceID,
		"appRoleId":   appRoleID,
	}

	u := msgraph.URL(Groups, groupID).Append(AppRoleAssignments).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, requestData)
	if err != nil {
		return nil, err
	}

	appRoleAssignmentResp := new(AppRoleAssignmentResponse)
	_, err = s.client.Do(ctx, req, appRoleAssignmentResp)
	if err != nil {
		return nil, err
	}

	return appRoleAssignmentResp, nil
}

// DeleteAppRoleAssignment deletes an appRoleAssignment that a group has been granted.
//
// DELETE /groups/{id}/appRoleAssignments/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-delete-approleassignments?view=graph-rest-beta&tabs=http
func (s *GroupsService) DeleteAppRoleAssignment(ctx context.Context, groupID, appRoleAssignmentID string) error {
	u := msgraph.URL(Groups, groupID).AppendWithID(AppRoleAssignments, appRoleAssignmentID).Build()
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
