package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	"github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

type UsersService service

const (
	UsersResourceName = "users"
)

type UserListOptions struct {
	// The $count and $search parameters are currently not available in Azure AD B2C tenants.
	//Search string `url:"$search,omitempty"`
	//Count  bool   `url:"$count,omitempty"`
	Filter   string `url:"$filter,omitempty"`
	Select   string `url:"$select,omitempty"`
	Top      int    `url:"$top,omitempty"`
	NextLink string `url:"-"`
}

type UsersResponse struct {
	models.OData
	Users []*models.User `json:"value"`
}

// Retrieve a list of user objects.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http
func (s *UsersService) ListAll(ctx context.Context, opts *UserListOptions) (*UsersResponse, error) {
	var u string
	if opts.NextLink != "" {
		u = opts.NextLink
	} else {
		var err error
		u, err = core.AddOptions(UsersResourceName, opts)
		if err != nil {
			return nil, err
		}
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
	models.OData
	models.User
}

func (u *UserResponse) UnmarshalJSON(data []byte) error {
	var usr models.User
	err := json.Unmarshal(data, &usr)
	if err != nil {
		return err
	}

	var odata models.OData
	if err := json.Unmarshal(data, &odata); err != nil {
		return err
	}
	u.OData = odata
	u.User = usr
	return nil
}

// Retrieve the properties and relationships of user object.
//
// GET /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-get?view=graph-rest-1.0&tabs=http
func (s *UsersService) GetUser(ctx context.Context, userID string, opts *UserListOptions) (*UserResponse, error) {
	u, err := core.AddOptions(fmt.Sprintf("users/%v", userID), opts)
	if err != nil {
		return nil, err
	}

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

// Create a new user. The request body contains the user to create.
// At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-post-users?view=graph-rest-1.0&tabs=http
func (s *UsersService) CreateUser(ctx context.Context, r *models.User) (*UserResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, UsersResourceName, r)
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

// Update the properties of a user object. Not all properties can be updated by Member or Guest users with their
// default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
//
// PATCH /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-update?view=graph-rest-1.0&tabs=http
func (s *UsersService) UpdateUser(ctx context.Context, u *models.User) error {
	req, err := s.client.NewRequest(http.MethodPatch, core.BuildURL(UsersResourceName, *u.Id), u)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete user. When deleted, user resources are moved to a temporary container and can be restored within 30 days.
// After that time, they are permanently deleted. To learn more, see deletedItems.
// At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties.
//
// DELETE /users/{id | userPrincipalName}
//
// MS Graph API doc:
// https://docs.microsoft.com/en-us/graph/api/user-delete?view=graph-rest-1.0&tabs=http
func (s *UsersService) DeleteUser(ctx context.Context, userID string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("users/%v", userID), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Update the schema extensions of a user object.
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
// https://docs.microsoft.com/en-us/graph/api/user-update?view=graph-rest-1.0&tabs=http
func (s *UsersService) UpdateSchemaExtensions(ctx context.Context, userID, schemaExtensionID string, schemaExtensionData map[string]string) error {
	schemaExtensionDataReq := &models.SchemaExtensionRawDataType{
		schemaExtensionID: schemaExtensionData,
	}

	req, err := s.client.NewRequest(http.MethodPatch, core.BuildURL(UsersResourceName, userID), schemaExtensionDataReq)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
