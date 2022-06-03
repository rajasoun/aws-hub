[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=rajasoun_aws-hub&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=rajasoun_aws-hub)[![CodeQL - Vulnerability Analysis](https://github.com/rajasoun/aws-hub/workflows/CodeQL/badge.svg)](https://github.com/rajasoun/aws-hub/actions?query=workflow:CodeQL)[![Dependency Review - Vulnerability Check](https://github.com/rajasoun/aws-hub/workflows/DependencyReview/badge.svg)](https://github.com/rajasoun/aws-hub/actions?query=workflow:DependencyReview)[![Secrets Scan](https://github.com/rajasoun/aws-hub/workflows/SecretScan/badge.svg)](https://github.com/rajasoun/aws-hub/actions?query=workflow:SecretScan)[![Go Build](https://github.com/rajasoun/aws-hub/workflows/GoBuild/badge.svg)](https://github.com/rajasoun/aws-hub/actions?query=workflow:GoBuild)


# AWS Cot Hub

Go based cost exploration tool for AWS on multiple accounts

## Getting Started

### 1. Setup

Refer [Setup Instructions](docs/setup.md)

### 2. Start the Server

1. In Visual Studio Code, Click the Green Button as shown in the image below and select.
   - Open Folder in Container... command and select the local folder.
   - Click the ![Green Button](docs/images/remote-status-bar.png)

1. Start the server

   ```bash
   $ go mod tidy
   $ aws-env go run main.go start
   ```

1. In another terminal, query the rest end point

   ```bash
   $ AWS_PROFILE=<profile_name>
   $ http localhost:3000/aws/iam/account 'profile:$AWS_PROFILE'
   ```

### 3. aws-identity from CLI

Get Identity by account - validate setup
   ```bash
   $ aws-whoami <aws_profile>
   ```

### 4. AWS Consolidated Bill

For previous month using aws-vault via aws cli

1. Get Identity by account. Select aws-vault
   ```bash
   $ /workspaces/tools/all_bills.sh
   ```

1. View Bill

   ```bash
   $ column -t -s,  /tmp/reports/aws-cost-hub/aws-cli/bill.csv
   ```

## Test Driven Development

### Server Start - Key Flow

1. Refer [Server Start Flow](docs/flow.md)

1. In Visual Studio Code devcontainer Terminal, Install Go Packages needed for TDD

   ```sh
   touch .dev
   make install-packages
   ```

1. Running All Tests

   ```sh
   make tdd-unit
   aws-env make tdd-integration
   ```

1. Running Package Level Tests

   ```sh
   gotestsum --format testname -- -coverprofile=coverage/coverage.out github.com/rajasoun/aws-hub/app/...
   ```
