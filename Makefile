APP := $(shell basename $(CURDIR))
VERSION := $(shell git describe --tags --always --dirty)
#GOPATH := $(CURDIR)/Godeps/_workspace:$(GOPATH)
#PATH := $(GOPATH)/bin:$(PATH)

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
#.PHONY: bin/$(APP) bin clean start test

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# Go TASKS
build:  build/bin/$(APP) test ## Build Go

build/bin/$(APP): bin
	go build -v -o $@ -ldflags "-X main.Version='${VERSION}'"

bin: clean
	mkdir -p build/bin

clean: ## Clean Go
	rm -rf build/bin

lint: ## Go Lint
	golangci-lint run --enable-all

tdd:  ## Test Go
	go test ./... -v

tdd-watch: ## Test Watch
	gotestsum --watch --format testname

tdd-cover: ## Go Coverage
	go test ./... -v --cover -coverprofile build/coverage/coverage.out
	go tool cover -html=coverage/coverage.out

tdd-unit: ## Prints formatted unit test output
	export SKIP_E2E=true && gotestsum --format testname -- -coverprofile=build/coverage/coverage.out ./...
	@bash -c "test/coverage_check.sh"

tdd-integration: ## Prints formatted integration test output
	gotestsum --format testname -- -coverprofile=build/coverage/coverage.out ./...
	@bash -c "test/coverage_check.sh"

install-packages: ## Install go packages
	go install -v gotest.tools/gotestsum@latest
	go install -v github.com/cweill/gotests/gotests@latest
