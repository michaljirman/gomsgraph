package beta

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

type UsersService service

type UsersResponse struct {
	OData
	Users []*User `json:"value"`
}

// ListAll retrieves a list of user objects.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-beta&tabs=http
func (s *UsersService) ListAll(ctx context.Context, opts *ListOptions) (*UsersResponse, error) {
	var u string
	if opts.NextLink != "" {
		u = opts.NextLink
	} else {
		u = msgraph.URL(Users).Options(opts).Build()
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var usersResponse UsersResponse
	_, err = s.client.Do(ctx, req, &usersResponse)
	if err != nil {
		return nil, err
	}

	return &usersResponse, nil
}

type UserResponse struct {
	OData
	User
}

func (u *UserResponse) UnmarshalJSON(data []byte) error {
	var usr User
	err := json.Unmarshal(data, &usr)
	if err != nil {
		return err
	}

	var odata OData
	if err := json.Unmarshal(data, &odata); err != nil {
		return err
	}
	u.OData = odata
	u.User = usr
	return nil
}

// GetUser retrieves the properties and relationships of user object.
//
// GET /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-get?view=graph-rest-beta&tabs=http
func (s *UsersService) GetUser(ctx context.Context, userID string, opts *ListOptions) (*UserResponse, error) {
	u := msgraph.URL(Users, userID).Options(opts).Build()
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	userResp := new(UserResponse)
	_, err = s.client.Do(ctx, req, userResp)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}

// CreateUser creates a new user. The request body contains the user to create.
// At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-post-users?view=graph-rest-beta&tabs=http
func (s *UsersService) CreateUser(ctx context.Context, r *User) (*UserResponse, error) {
	u := msgraph.URL(Users).Build()
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}

	userResp := new(UserResponse)
	_, err = s.client.Do(ctx, req, userResp)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}

// UpdateUser updates the properties of a user object. Not all properties can be updated by Member or Guest users with their
// default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
//
// PATCH /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-update?view=graph-rest-beta&tabs=http
func (s *UsersService) UpdateUser(ctx context.Context, usr User) error {
	if usr.Id == nil {
		return errors.New("user id is required")
	}
	u := msgraph.URL(Users, *usr.Id).Build()
	req, err := s.client.NewRequest(http.MethodPatch, u, usr)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes user. When deleted, user resources are moved to a temporary container and can be restored within 30 days.
// After that time, they are permanently deleted. To learn more, see deletedItems.
// At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
//
// DELETE /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-delete?view=graph-rest-beta&tabs=http
func (s *UsersService) DeleteUser(ctx context.Context, userID string) error {
	u := msgraph.URL(Users, userID).Build()
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

// UpdateSchemaExtensions updates the schema extensions of a user object.
// You can extend and add custom data to an existing group instance with the additional graphlearn_courses complex type extension defined in the body of a PATCH request.
// If you want to update the values of the extension data, put the entire extension complex type in the body of a PATCH request (similar to adding custom data to an existing resource).
// You can also remove custom data added to a resource instance by setting the corresponding extension property to null.
// To remove a schema extension from a resource instance, set the extension complex type in that instance to null.
//
// PATCH /users/{id | userPrincipalName}
//
// Warning:
// * One application can only create up to 5 schema extensions. Keep that in mind!
// * Extension schema can be deleted only if the status is not set to "Available". (InDevelopment=>Available=>Deprecated)
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-update?view=graph-rest-beta&tabs=http
func (s *UsersService) UpdateSchemaExtensions(ctx context.Context, userID, schemaExtensionID string, schemaExtensionData map[string]string) error {
	schemaExtensionDataReq := &SchemaExtensionRawDataType{
		schemaExtensionID: schemaExtensionData,
	}
	u := msgraph.URL(Users, userID).Build()
	req, err := s.client.NewRequest(http.MethodPatch, u, schemaExtensionDataReq)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListAppRoleAssignments retrieves the list of appRoleAssignment that a user has been granted.
// This operation also returns app roles assigned to groups that the user is a direct member of
//
// GET /users/{id | userPrincipalName}/appRoleAssignments
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-list-approleassignments?view=graph-rest-beta&tabs=http
func (s *UsersService) ListAppRoleAssignments(ctx context.Context, userID string, opts *ListOptions) (*AppRoleAssignmentsResponse, error) {
	u := msgraph.URL(Users, userID).Append(AppRoleAssignments).Options(opts).Build()
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

// AddAppRoleAssignments assigns an app role to a user. To grant an app role assignment to a user, you need three identifiers:
//
// principalId: The id of the group to which you are assigning the app role.
// resourceId: The id of the resource servicePrincipal which has defined the app role.
// appRoleId: The id of the appRole (defined on the resource service principal) to assign to the group.
//
// POST /users/{id | userPrincipalName}/appRoleAssignments
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-post-approleassignments?view=graph-rest-beta&tabs=http
func (s *UsersService) AddAppRoleAssignments(ctx context.Context, userID, principalID, resourceID, appRoleID string) (*AppRoleAssignmentResponse, error) {
	requestData := map[string]string{
		"principalId": principalID,
		"resourceId":  resourceID,
		"appRoleId":   appRoleID,
	}

	u := msgraph.URL(Users, userID).Append(AppRoleAssignments).Build()
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

// DeleteAppRoleAssignment deletes an appRoleAssignment that has been granted to a user.
//
// DELETE /groups/{id}/appRoleAssignments/{id}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/group-delete-approleassignments?view=graph-rest-beta&tabs=http
func (s *UsersService) DeleteAppRoleAssignment(ctx context.Context, userID, appRoleAssignmentID string) error {
	u := msgraph.URL(Users, userID).AppendWithID(AppRoleAssignments, appRoleAssignmentID).Build()
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
