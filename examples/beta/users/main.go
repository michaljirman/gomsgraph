package main

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/michaljirman/gomsgraph/msgraph"
	"github.com/michaljirman/gomsgraph/msgraph/beta"
	. "github.com/michaljirman/gomsgraph/msgraph/beta/models"
)

func main() {
	ctx := context.Background()
	client := beta.NewDefaultClient(ctx)

	req := &User{
		Id:                msgraph.String(uuid.NewString()),
		AccountEnabled:    msgraph.Bool(true),
		DisplayName:       msgraph.String("Test User"),
		MailNickname:      msgraph.String("TestU"),
		UserPrincipalName: msgraph.String("TestU@testworkspace.onmicrosoft.com"),
		PasswordProfile: &PasswordProfile{
			Password:                             msgraph.String(uuid.NewString()),
			ForceChangePasswordNextSignIn:        msgraph.Bool(true),
			ForceChangePasswordNextSignInWithMfa: msgraph.Bool(true),
		},
	}
	userResp, err := client.Users.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("failed to create a new user %v", err)
	}

	usersResp, err := client.Users.ListAll(ctx, &ListOptions{})
	if err != nil {
		log.Fatalf("failed to list all users %v", err)
	}
	log.Printf("%d users found\n", len(usersResp.Users))

	user, err := client.Users.GetUser(ctx, *userResp.Id, nil)
	if err != nil {
		log.Fatalf("failed to get user by ID %v", err)
	}

	log.Printf("user: %+v, %+v", user.Id, user.DisplayName)
}
