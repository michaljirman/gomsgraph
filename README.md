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
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/michaljirman/gomsgraph/msgraph"
	"github.com/michaljirman/gomsgraph/msgraph/models"
)

var (
	azureClientID     string
	azureClientSecret string
	azureTenantID     string
)

func main() {
	flag.StringVar(&azureClientID, "azure-client-id", "", "")
	flag.StringVar(&azureClientSecret, "azure-client-secret", "", "")
	flag.StringVar(&azureTenantID, "azure-client-tenant-id", "", "")
	flag.Parse()

	if azureClientID == "" {
		log.Fatal("azure client id is required")
	}

	if azureClientSecret == "" {
		log.Fatal("azure client secret is required")
	}

	if azureTenantID == "" {
		log.Fatal("azure tenant id is required")
	}

	ctx := context.Background()
	tc := msgraph.GetOAuth2Client(ctx, azureTenantID, azureClientID, azureClientSecret)
	client := msgraph.NewClient(tc)
}
```

To create a new User: 
```go
ctx := context.Background()
req := &models.User{
    Id:                msgraph.String(uuid.NewString()),
    AccountEnabled:    msgraph.Bool(true),
    DisplayName:       msgraph.String("Test User"),
    MailNickname:      msgraph.String("TestU"),
    UserPrincipalName: msgraph.String("TestU@testworkspace.onmicrosoft.com"),
    PasswordProfile: &models.PasswordProfile{
        Password:                             msgraph.String(uuid.NewString()),
        ForceChangePasswordNextSignIn:        msgraph.Bool(true),
        ForceChangePasswordNextSignInWithMfa: msgraph.Bool(true),
    },
}
userResp, err := client.Users.CreateUser(ctx, req)
```

To delete a user:
```go
err = client.Users.DeleteUser(ctx, *userResp.Id)
```

To list all users:
```go
usersResp, err := client.Users.ListAll(ctx, &msgraph.UserListOptions{})
```

To list users with pagination:
```go
opts := &msgraph.UserListOptions{
    Top: 1,
}
var allUsers []*models.User
for {
    usersResp, err := client.Users.ListAll(ctx, opts)
    allUsers = append(allUsers, usersResp.Users...)
    if usersResp.NextLink == nil {
        break
    }
    opts.NextLink = *usersResp.NextLink
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

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).



