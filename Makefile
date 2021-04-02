.PHONY: fmt
## fmt              		: format go files
fmt:
	@echo "--> Formatting go files"
	@go fmt $$(go list ./...)

.PHONY: golangci-lint
## golangci-lint            	: run golangci-lint
golangci-lint:
	cd /tmp && GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: staticcheck
## staticcheck              	: run staticcheck
staticcheck:
	cd /tmp && GO111MODULE=on go get honnef.co/go/tools/cmd/staticcheck@master
	staticcheck ./...

.PHONY: lint
## lint              		: run lint tools
lint: golangci-lint staticcheck

.PHONY: pretest
pretest: lint

.PHONY: gotest
## gotest             		: run all golang tests
gotest:
	go test --race ./...

.PHONY: test
## test             		: run all tests
test: pretest gotest

.PHONY: integration
## integration             	: run all integration tests
integration:
	go test -tags integration ./test/integration/...

.PHONY: help
help: Makefile
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | sort