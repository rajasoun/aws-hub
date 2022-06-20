APP := $(shell basename $(CURDIR))
VERSION := $(shell git describe --tags --always --dirty)
IMAGE_NAME=rajasoun/$(APP):$(VERSION)
CONTEXT=.ci/container/
REPORTS_DIR=.ci/reports

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
build:  .ci/build/bin/$(APP) test ## Build Go

.ci/build/bin/$(APP): bin
	go build -v -installsuffix 'static' -o $@ -ldflags "-X main.Version='${VERSION}'"

check-for-updates:	## View minor/patch upgrades 
	go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null

bin: clean
	mkdir -p .ci/build/bin

clean: ## Clean Go
	rm -rf .ci/build/bin

lint: ## Go Lint
	golangci-lint run 

gosec: ## Lint Go Code for security issues
	@gosec -exclude=G104 -fmt=json -out=$(REPORTS_DIR)/security/results.json -stdout --verbose=text  ./...

tdd-unit: ## Prints formatted unit test output
	@export SKIP_E2E=true && gotestsum --format testname -- -coverprofile=$(REPORTS_DIR)/coverage/coverage.out ./...
	@go tool cover -html=$(REPORTS_DIR)/coverage/coverage.out -o $(REPORTS_DIR)/coverage/coverage.html
	@bash -c "test/coverage_check.sh"

tdd-integration: ## Prints formatted integration test output
	gotestsum --format testname -- test/integration/**.go

tdd-e2e: ## Prints formatted e2e test output
	gotestsum --format testname -- test/e2e/api_test.go
	gotestsum --format testname -- test/e2e/understand_test.go
	
docker-build: ## Build aws-hub docker container
	docker build  -t $(IMAGE_NAME) $(CONTEXT) 

# docker-start: ## Run container 
# 	docker run --rm --name $(APP) --publish 3000:3000 -v "${PWD}:/workspace" $(IMAGE_NAME)

# docker-stop: ## Stop container 
# 	docker stop $(APP)

check-all: tdd-unit lint ## Check Lint & Unit Test

install-packages: ## Install go packages
	go install -v golang.org/x/tools/gopls@latest
	go install -v gotest.tools/gotestsum@latest
	go install -v github.com/cweill/gotests/gotests@latest
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -v github.com/securego/gosec/v2/cmd/gosec@latest
