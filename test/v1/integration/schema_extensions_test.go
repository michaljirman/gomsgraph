// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/michaljirman/gomsgraph/msgraph"
	. "github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func TestListSchemaExtensionsWithPagination(t *testing.T) {
	ctx := context.Background()
	opts := &ListOptions{}
	var allSchemaExtensions []*SchemaExtension
	for {
		resp, err := client.SchemaExtensions.ListAll(ctx, opts)
		require.NoError(t, err)
		allSchemaExtensions = append(allSchemaExtensions, resp.SchemaExtensions...)
		if resp.NextLink == nil {
			break
		}
		opts.NextLink = *resp.NextLink
	}
}

func TestCreateSchemaExtension(t *testing.T) {
	// create a test user
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

	// create a test schema extension for a test user
	schemaExtResource := SchemaExtension{
		Id:          msgraph.String("courses"),
		Description: msgraph.String("Graph Learn training courses extensions"),
		TargetTypes: &[]string{
			"User",
		},
		Properties: &[]ExtensionSchemaProperty{
			{
				Name: msgraph.String("courseName"),
				Type: msgraph.String("String"),
			},
		},
	}
	schemaExtResp, err := client.SchemaExtensions.CreateSchemaExtension(ctx, schemaExtResource)
	require.NoError(t, err)

	defer func() {
		err := client.SchemaExtensions.DeleteSchemaExtension(ctx, *schemaExtResp.Id)
		require.NoError(t, err)
	}()

	schemaExtData := map[string]string{
		"courseName": "test-course-name",
	}

	require.Eventually(t, func() bool {
		// it takes time before a newly created extension can be used
		err := client.Users.UpdateSchemaExtensions(ctx, *userResp.Id, *schemaExtResp.Id, schemaExtData)
		return err == nil
	}, 10*time.Second, 2*time.Second)

	require.Eventually(t, func() bool {
		// fetch user data with schema extension properties stored in the AdditionalData
		getUserResp, err := client.Users.GetUser(ctx, *userResp.Id, &ListOptions{
			Select: "displayName,id," + *schemaExtResp.Id,
		})
		require.NoError(t, err)
		// check if the AdditionalData contains schema extension properties
		_, schemaExtDataExistsInResp := getUserResp.AdditionalData[*schemaExtResp.Id]
		return schemaExtDataExistsInResp
	}, 20*time.Second, 4*time.Second)
}
