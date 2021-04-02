// +build integration

package v1

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	msgraph "github.com/michaljirman/gomsgraph/msgraph/v1"
	"github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func TestUsers(t *testing.T) {
	// create a user
	ctx := context.Background()
	req := &models.User{
		Id:                core.String(uuid.NewString()),
		AccountEnabled:    core.Bool(true),
		DisplayName:       core.String("Test User"),
		MailNickname:      core.String("TestU"),
		UserPrincipalName: core.String("TestU@testebworkspace.onmicrosoft.com"),
		PasswordProfile: &models.PasswordProfile{
			Password:                             core.String(uuid.NewString()),
			ForceChangePasswordNextSignIn:        core.Bool(true),
			ForceChangePasswordNextSignInWithMfa: core.Bool(true),
		},
	}
	userResp, err := client.Users.CreateUser(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, userResp)

	defer func() {
		// delete a user
		err = client.Users.DeleteUser(ctx, *userResp.Id)
		require.NoError(t, err)
	}()

	require.Eventually(t, func() bool {
		// list all users
		usersResp, err := client.Users.ListAll(ctx, &msgraph.UserListOptions{})
		if err != nil {
			return false
		}

		return len(usersResp.Users) > 0
	}, 5*time.Second, 1*time.Second)

	require.Eventually(t, func() bool {
		// list all users matching filter
		usersResp, err := client.Users.ListAll(ctx, &msgraph.UserListOptions{
			Filter: "startswith(displayName,'Test')",
		})
		if err != nil {
			return false
		}

		return len(usersResp.Users) > 0
	}, 5*time.Second, 1*time.Second)

	// get a user
	getUserResp, err := client.Users.GetUser(ctx, *userResp.Id, nil)
	require.NoError(t, err)
	require.NotNil(t, getUserResp)
}

func TestUsersWithPagination(t *testing.T) {
	// create a user
	ctx := context.Background()

	requiredNumberOfTestUsers := 5
	testUserIDs := make([]string, 0, requiredNumberOfTestUsers)
	for i := 0; i < requiredNumberOfTestUsers; i++ {
		req := &models.User{
			Id:                core.String(uuid.NewString()),
			AccountEnabled:    core.Bool(true),
			DisplayName:       core.String(fmt.Sprintf("Test User%s", strconv.Itoa(i))),
			MailNickname:      core.String(fmt.Sprintf("TestU%s", strconv.Itoa(i))),
			UserPrincipalName: core.String(fmt.Sprintf("TestU%s@testebworkspace.onmicrosoft.com", strconv.Itoa(i))),
			PasswordProfile: &models.PasswordProfile{
				Password:                             core.String(uuid.NewString()),
				ForceChangePasswordNextSignIn:        core.Bool(true),
				ForceChangePasswordNextSignInWithMfa: core.Bool(true),
			},
		}
		userResp, err := client.Users.CreateUser(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, userResp)
		testUserIDs = append(testUserIDs, *userResp.Id)
	}

	defer func() {
		// delete a user
		for _, userID := range testUserIDs {
			err := client.Users.DeleteUser(ctx, userID)
			require.NoError(t, err)
		}
	}()

	// list all users
	opts := &msgraph.UserListOptions{
		Top: 1,
	}
	var allUsers []*models.User
	for {
		usersResp, err := client.Users.ListAll(ctx, opts)
		require.NoError(t, err)
		allUsers = append(allUsers, usersResp.Users...)
		fmt.Printf("page with size %d\n", len(usersResp.Users))
		if usersResp.NextLink == nil {
			break
		}
		opts.NextLink = *usersResp.NextLink
	}
	require.GreaterOrEqual(t, len(allUsers), requiredNumberOfTestUsers)
}
