// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func TestGroups(t *testing.T) {
	// create a group
	ctx := context.Background()
	req := Group{
		Description: msgraph.String("Self help community for library"),
		DisplayName: msgraph.String("Library Assist"),
		GroupTypes: &[]string{
			"Unified",
		},
		MailEnabled:     msgraph.Bool(true),
		MailNickname:    msgraph.String("library"),
		SecurityEnabled: msgraph.Bool(false),
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
		groupsResp, err := client.Groups.ListAll(ctx, &ListOptions{})
		if err != nil {
			return false
		}

		return len(groupsResp.Groups) > 0
	}, 5*time.Second, 1*time.Second)

	require.Eventually(t, func() bool {
		// list all groups matching filter
		groupsResp, err := client.Groups.ListAll(ctx, &ListOptions{
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
