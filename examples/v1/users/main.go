package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/michaljirman/gomsgraph/msgraph/core"
	msgraph "github.com/michaljirman/gomsgraph/msgraph/v1"
	"github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func main() {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("env variable CLIENT_ID is required for integration tests")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("env variable CLIENT_SECRET is required for integration tests")
	}

	clientTenantID := os.Getenv("CLIENT_TENANT_ID")
	if clientTenantID == "" {
		log.Fatal("env variable CLIENT_TENANT_ID is required for integration tests")
	}

	ctx := context.Background()
	tc := core.GetOAuth2Client(ctx, clientTenantID, clientID, clientSecret)
	client := msgraph.NewClient(tc)

	req := &models.User{
		Id:                core.String(uuid.NewString()),
		AccountEnabled:    core.Bool(true),
		DisplayName:       core.String("Test User"),
		MailNickname:      core.String("TestU"),
		UserPrincipalName: core.String("TestU@testworkspace.onmicrosoft.com"),
		PasswordProfile: &models.PasswordProfile{
			Password:                             core.String(uuid.NewString()),
			ForceChangePasswordNextSignIn:        core.Bool(true),
			ForceChangePasswordNextSignInWithMfa: core.Bool(true),
		},
	}
	userResp, err := client.Users.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("failed to create a new user %v", err)
	}

	usersResp, err := client.Users.ListAll(ctx, &msgraph.UserListOptions{})
	if err != nil {
		log.Fatalf("failed to list all users %v", err)
	}
	fmt.Printf("%d users found", len(usersResp.Users))

	user, err := client.Users.GetUser(ctx, *userResp.Id, nil)
	if err != nil {
		log.Fatalf("failed to get user by ID %v", err)
	}
	fmt.Println(user)
}
