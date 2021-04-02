// +build integration

package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph/beta/models"
	"github.com/michaljirman/gomsgraph/msgraph/core"
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

	req := &models.Group{
		Description: core.String("Self help community for library"),
		DisplayName: core.String("Library Assist"),
		GroupTypes: &[]string{
			"Unified",
		},
		MailEnabled:        core.Bool(true),
		MailNickname:       core.String("library"),
		SecurityEnabled:    core.Bool(true),
		IsAssignableToRole: core.Bool(true),
		Visibility:         core.String("Private"),
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
