// +build integration

package integration

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func TestUsers(t *testing.T) {
	// create a user
	ctx := context.Background()
	req := User{
		Id:                msgraph.String(uuid.NewString()),
		AccountEnabled:    msgraph.Bool(true),
		DisplayName:       msgraph.String("Test User"),
		MailNickname:      msgraph.String("TestU"),
		UserPrincipalName: msgraph.String("TestU@testebworkspace.onmicrosoft.com"),
		PasswordProfile: &PasswordProfile{
			Password:                             msgraph.String(uuid.NewString()),
			ForceChangePasswordNextSignIn:        msgraph.Bool(true),
			ForceChangePasswordNextSignInWithMfa: msgraph.Bool(true),
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
		usersResp, err := client.Users.ListAll(ctx, &ListOptions{})
		if err != nil {
			return false
		}

		return len(usersResp.Users) > 0
	}, 5*time.Second, 1*time.Second)

	require.Eventually(t, func() bool {
		// list all users matching filter
		usersResp, err := client.Users.ListAll(ctx, &ListOptions{
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
		req := User{
			Id:                msgraph.String(uuid.NewString()),
			AccountEnabled:    msgraph.Bool(true),
			DisplayName:       msgraph.String(fmt.Sprintf("Test User%s", strconv.Itoa(i))),
			MailNickname:      msgraph.String(fmt.Sprintf("TestU%s", strconv.Itoa(i))),
			UserPrincipalName: msgraph.String(fmt.Sprintf("TestU%s@testebworkspace.onmicrosoft.com", strconv.Itoa(i))),
			PasswordProfile: &PasswordProfile{
				Password:                             msgraph.String(uuid.NewString()),
				ForceChangePasswordNextSignIn:        msgraph.Bool(true),
				ForceChangePasswordNextSignInWithMfa: msgraph.Bool(true),
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
	opts := &ListOptions{
		Top: 1,
	}
	var allUsers []*User
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
