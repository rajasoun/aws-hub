#!/usr/bin/env bash

function fix_n_commits_behind(){
    main_branch=$1
    dev_branch=$2
    git checkout $main_branch
    git pull origin $main_branch --rebase
    git checkout $dev_branch
    git merge $main_branch
    # dev_branch is in sync with local main_branch
    git push origin dev_branch
}

fix_n_commits_behind "main" "develop"
fix_n_commits_behind "develop" "integration_branch"