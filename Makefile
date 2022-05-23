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
build:  bin/$(APP) test ## Build Go

bin/$(APP): bin
	go build -v -o $@ -ldflags "-X main.Version='${VERSION}'"

bin: clean
	mkdir -p bin

clean: ## Clean Go
	rm -rf bin

lint: ## Go Lint
	golangci-lint run --enable-all

tdd:  ## Test Go
	go test ./... -v

tdd-watch: ## Test Watch
	gotestsum --watch --format testname

tdd-cover: ## Go Coverage
	go test ./... -v --cover -coverprofile coverage/coverage.out
	go tool cover -html=coverage/coverage.out

tdd-summary: ## Prints formatted test output
	gotestsum --format testname -- -coverprofile=coverage/coverage.out ./...

install-packages: ## Install go packages
	go install -v gotest.tools/gotestsum@latest
	go install -v github.com/cweill/gotests/gotests@latest
