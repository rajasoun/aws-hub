#!/usr/bin/env bash

# Means your copy of the remote main branch (a.ka.a origin/main) 
# has n commits more than the local version of the main branch. 
# Below Function fixes teh issue
# Typically happens is fast paced development environment with long running branches
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
    fix_n_commits_behind "main" "integration_branch"
}

main

