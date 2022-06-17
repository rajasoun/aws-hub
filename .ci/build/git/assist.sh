#!/usr/bin/env bash

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

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

# Remove Branches that are not in remote
# Remove branches that are already synced
function clean_branches(){
    # prunes tracking branches not on the remote
    git remote prune origin
    # ignore main branch and delete other merged branch
    git branch --merged | grep -v "\*" | grep -v "main" | xargs -n 1 --no-run-if-empty  git branch -d
    git gc 
}

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )

case ${choice} in
    "clean")
        clean_branches
    ;;
    "sync")
        current_branch=$( git rev-parse --abbrev-ref HEAD)
        fix_n_commits_behind "main" $current_branch
    ;;
    *)
    echo -e "\n${RED}Usage: build/git/assist.sh < clean | sync >${NC}\n"
cat <<-EOF
Commands:
---------
  clean       -> Remove Branches Not In Remote or Merged
  sync        -> Sync Current Branch with Main (for shared branches)

EOF
    ;;
esac

