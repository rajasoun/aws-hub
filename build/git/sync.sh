#!/usr/bin/env bash

function fix_n_commits_behind(){
    upstream_branch=$1
    current_branch=$2
    git checkout $upstream_branch
    git pull origin $upstream_branch --rebase
    git checkout $current_branch
    git merge $upstream_branch
    # dev_branch is in sync with local main_branch
    git push origin $current_branch
}

function main(){
    fix_n_commits_behind "develop" "integration_branch"
    fix_n_commits_behind "main" "develop"
}

main

