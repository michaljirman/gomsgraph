// +build integration

package v1

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	msgraph "github.com/michaljirman/gomsgraph/msgraph/v1"
	"github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func TestGroups(t *testing.T) {
	// create a group
	ctx := context.Background()
	req := &models.Group{
		Description: core.String("Self help community for library"),
		DisplayName: core.String("Library Assist"),
		GroupTypes: &[]string{
			"Unified",
		},
		MailEnabled:     core.Bool(true),
		MailNickname:    core.String("library"),
		SecurityEnabled: core.Bool(false),
	}
	groupResp, err := client.Groups.CreateGroup(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, groupResp)

	defer func() {
		// delete a group
		err = client.Groups.DeleteGroup(ctx, *groupResp.Id)
		require.NoError(t, err)
	}()

	require.Eventually(t, func() bool {
		// list all groups
		groupsResp, err := client.Groups.ListAll(ctx, &msgraph.GroupListOptions{})
		if err != nil {
			return false
		}

		return len(groupsResp.Groups) > 0
	}, 5*time.Second, 1*time.Second)

	require.Eventually(t, func() bool {
		// list all groups matching filter
		groupsResp, err := client.Groups.ListAll(ctx, &msgraph.GroupListOptions{
			Filter: "startswith(displayName,'Library')",
		})
		if err != nil {
			return false
		}

		return len(groupsResp.Groups) > 0
	}, 5*time.Second, 1*time.Second)

	// get a group
	getGroupResp, err := client.Groups.GetGroup(ctx, *groupResp.Id, nil)
	require.NoError(t, err)
	require.NotNil(t, getGroupResp)
}
