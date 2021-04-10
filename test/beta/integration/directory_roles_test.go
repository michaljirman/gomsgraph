// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

func TestDirectoryRoles(t *testing.T) {
	ctx := context.Background()
	// list all directory roles
	dirRolesResp, err := client.DirectoryRoles.ListAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, dirRolesResp)
	require.NotEmpty(t, dirRolesResp.DirectoryRoles)

	// pick id of first directory role in the list
	dirRoleID := *dirRolesResp.DirectoryRoles[0].Id

	// get a directory role
	dirRoleResp, err := client.DirectoryRoles.GetDirectoryRole(ctx, dirRoleID)
	require.NoError(t, err)
	require.NotNil(t, dirRoleResp)

	req := &Group{
		Description: msgraph.String("Self help community for library"),
		DisplayName: msgraph.String("Library Assist"),
		GroupTypes: &[]string{
			"Unified",
		},
		MailEnabled:     msgraph.Bool(true),
		MailNickname:    msgraph.String("library"),
		SecurityEnabled: msgraph.Bool(true),
		// set IsAssignableToRole to allow roles assignment to a newly created group
		IsAssignableToRole: msgraph.Bool(true),
		Visibility:         msgraph.String("Private"),
	}
	groupResp, err := client.Groups.CreateGroup(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, groupResp)

	defer func() {
		// delete a group
		err = client.Groups.DeleteGroup(ctx, *groupResp.Id)
		require.NoError(t, err)
	}()

	// Assign a directory role to a group
	err = client.DirectoryRoles.AddMember(ctx, dirRoleID, *groupResp.Id)
	require.NoError(t, err)
}
