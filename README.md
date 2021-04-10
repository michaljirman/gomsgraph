# gomsgraph #

[![gomsgraph release (latest SemVer)](https://img.shields.io/github/v/release/michaljirman/gomsgraph?sort=semver)](https://github.com/michaljirman/gomsgraph/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaljirman/gomsgraph/msgraph)
[![Test Status](https://github.com/michaljirman/gomsgraph/workflows/tests/badge.svg)](https://github.com/michaljirman/gomsgraph/actions?query=workflow%3Atests)

Gomsgraph is a Go client library for accessing the [Microsoft Graph API v1](https://docs.microsoft.com/en-us/graph/overview?view=graph-rest-1.0).

## Package documentation

For complete usage of the Gomsgraph client, see the full [package documentation](https://pkg.go.dev/github.com/michaljirman/gomsgraph).

## Installation ##

Installation of a specific [version](https://github.com/michaljirman/gomsgraph/releases) using `go get`.
```bash
go get github.com/michaljirman/gomsgraph@v1.0.0
```

## Authentication ##
The Gomsgraph library uses [oauth2](https://github.com/golang/oauth2) library to implement the oauth2 authentication. 
Simply pass `http.Client` implementation which can handle authentication while creating a new `MS Graph` client.   

```go
ctx := context.Background()
tc := msgraph.GetOAuth2Client(ctx, azureTenantID, azureClientID, azureClientSecret)
client := msgraph.NewClient(tc)
```

OAuth2 documentation is available on [oauth2 docs](https://godoc.org/golang.org/x/oauth2).

## Example usage ##

To create a new Gomsgraph client:
```go
package main

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/michaljirman/gomsgraph/msgraph"
	v1 "github.com/michaljirman/gomsgraph/msgraph/v1"
	. "github.com/michaljirman/gomsgraph/msgraph/v1/models"
)

func main() {
	ctx := context.Background()
	client := v1.NewDefaultClient(ctx)

	req := User{
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
```


## Developing ##
Committed code must pass following checks:

* [golangci-lint](https://github.com/golangci/golangci-lint)
* [go test](https://golang.org/cmd/go/#hdr-Test_packages)

To list all development commands:
```shell script
$ make help
 fmt              		: format go files
 golangci-lint            	: run golangci-lint
 gotest             		: run all golang tests
 integration             	: run all integration tests
 lint              		: run lint tools
 staticcheck              	: run staticcheck
 test             		: run all tests
```


To use a proxy, set the environment variable `HTTP_PROXY` from a shell or withing the program:
```bash
$ export HTTP_PROXY=http://proxy_name:proxy_port
```

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Versioning

Versions of the client are tagged on github. For each tagged version a new release is automatically creating 
using a [goreleaser](https://goreleaser.com/) and [github actions](https://github.com/features/actions).


## Contributing

Please see the [contribution guidelines](CONTRIBUTING.md).



