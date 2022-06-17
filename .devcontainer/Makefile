CONTEXT="."

BOLD :=$(shell tput bold)
RED :=$(shell tput setaf 1)
GREEN :=$(shell tput setaf 2)
YELLOW :=$(shell tput setaf 3)
RESET :=$(shell tput sgr0)

NAME := "aws-hub"
LIST_TAGS := $(shell git tag -l)
IMAGE_NAME=rajasoun/$(NAME)
VERSION := "1.0.0"

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# DOCKER TASKS
# Build the container

config-prerequisite: ## Pre-Requisite Checks for Configs
	@echo "$(GREEN)\nConfigs PreRequisite Checks $(RESET)"
	@ls -asl .devcontainer/dotfiles/.gitconfig >/dev/null 2>&1 && (echo "✅ .gitconfig File. Using Command -> generate_git_config") || (echo  "❌ .gitconfig File. Execute -> generate_git_config")
	@ls -asl .devcontainer/.ssh/*id_rsa* >/dev/null 2>&1 && (echo "✅ .ssh keys. Using Command -> gssh_config") || (echo  "❌ .ssh keys. Execute -> gssh_config")
	@ls -asl .devcontainer/.gpg2/keys/private.key >/dev/null 2>&1 && (echo "✅ gpg2 config. Using Command -> generate_gpg_keys") || (echo  "❌ gpg2 config. Execute -> generate_gpg_keys")
	@ls -asl .devcontainer/.store/.gpg-id >/dev/null 2>&1 && (echo "✅ password-store config. Using Command ->  init_pass_store") || (echo  "❌ password-store config. Execute -> init_pass_store")
	@ls -asl .devcontainer/.aws/config >/dev/null 2>&1 && (echo "✅ aws config. Using Command -> aws-env") || (echo  "❌ aws config. Execute -> aws-env")
	@echo ""

clean: ## Remove Base Container Image
	@echo "$(BOLD)$(YELLOW)Removing docker image locally - $(IMAGE_NAME) $(RESET)"
	$(eval IMAGES=$(shell docker images $(IMAGE_NAME) -a -q ))
	docker rmi --force $(IMAGES)
	@echo "$(BOLD)$(GREEN)Completed removing docker image - $(IMAGE_NAME) $(RESET)"

all:  prerequisite git build login push clean ## All in One Command - (prerequisite git build login push clean)
