# AWS Cot Hub

Go based cost exploration tool for AWS on multiple accounts

## Pre-Requisites

Windows Laptop - Refer to [win10x-onboard](https://github.com/rajasoun/win10x-onboard/blob/main/README.md)

Mac Laptop - Refer to [mac-onboard](https://github.com/rajasoun/mac-onboard/blob/main/README.md)

### 1. Clean Start

#### Windows

```sh
cd ~\workspace\win10x-onboard/
git pull --rebase
.\e2e.ps1 bash-it
nix/assist.sh bash-setup
nix/assist.sh clean
nix/assist.sh speed-test
cd ~/workspace
git clone https://github.com/rajasoun/aws-hub
cd aws-hub
git checkout develop
git pull --rebase
touch .dev
code .
```

#### macOS

```sh
wget -O- -q https://raw.githubusercontent.com/rajasoun/aws-toolz/main/all-in-one/speed.sh | bash
cd ~/workspace
git clone https://github.com/rajasoun/aws-hub
cd aws-hub
git checkout develop
git pull --rebase
touch .dev
code .
```


## 2. Development Setup

1.  In Visual Studio Code, Click the Green Button as shown in the image below and select
    Open Folder in Container... command and select the local folder.

    ![Click the Green Button](docs/images/remote-status-bar.png)

1.  Grab a Coffee ☕️. Based on your internet speed might take 5 mins to 7 mins with ~50 mbps speed

1. Open Terminal in visual studio code and run `/workspaces/tests/system/e2e_tests.sh ` for automated test of the setup

1. Run `ci-cd config-prerequisite` to configure gpg, pass and aws-vault for AWS Programatic access via API

### 3. Configure gpg, pass and aws-vault

1. Generate a new GPG private key. (Optional if you already have a GPG key setup and trusted on the system)
   > Note: If you set a passphrase, you will be prompted to enter it.

   ```bash
   $ generate_gpg_keys
   ```

1. Initialize the password-storage DB using the GPG `public` key ID or the associated email
   ```bash
   $ gpg2 --list-keys
   $ init_pass_store #similar to pass init <email_id> got from previous command
   ```
1. Configure aws-vault through wrapper
   ```bash
   $ aws-env
   ```

## 4. Start the Server

1. In Visual Studio Terminal , start the server

   ```bash
   $ go mod tidy
   $ aws-env go run main.go start
   ```

1. In another terminal, query the rest end point

   ```bash
   $ AWS_PROFILE=<profile_name>
   $ http localhost:3000/aws/iam/account 'profile:$AWS_PROFILE'
   ```

## 5. aws-identity

Get Identity by account
   ```bash
   $ aws-whoami <aws_profile>
   ```

## 6. AWS Consolidated Bill for Current month using aws-vault

1. Get Identity by account. Select aws-vault
   ```bash
   $ /workspaces/tools/all_bills.sh
   ```

1. View Bill

   ```bash
   $ column -t -s,  /tmp/reports/aws-hub/aws-cli/bill.csv
   ```

## Reset Configurations

To reset all configurations

```
clean_configs
ci-cd config-prerequisite
generate_git_config
gssh_config
generate_gpg_keys
gpg2 --list-keys
init_pass_store
aws-env
ci-cd config-prerequisite
```

## TDD

1. Install Go Packages needed for TDD
```sh
make install-packages
```

## Reference:

1. Generating Unit Tests
   ```
   gotests -w -only ^UnderstandStringReplace$ /workspaces/aws-hub/providers/env.go
   ```
