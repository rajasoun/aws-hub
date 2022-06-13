#!/usr/bin/env bash 

FAILED=()

function check_status(){
    echo -e "Checking $server/$url "
    url=$1 
    if http --check-status --ignore-stdin --timeout=5 HEAD $url 'profile:secops-experiments' &> /dev/null; then
        echo 'OK!'
    else
        FAILED+=("$url")
        case $? in
            2) echo 'Request timed out!' ;;
            3) echo 'Unexpected HTTP 3xx Redirection!' ;;
            4) echo 'HTTP 4xx Client Error!' ;;
            5) echo 'HTTP 5xx Server Error!' ;;
            6) echo 'Exceeded --max-redirects=<n> redirects!' ;;
            *) echo 'Other Error!' ;;
        esac
    fi
}

function reportResults() {
    failed_results="${#FAILED[@]}"
    if [ "${failed_results}" -ne 0 ]; then
        echoStderr -e "\n💥  Failed tests:" "${FAILED[@]}"
        exit 1
    else
        echo -e "\n💯  All passed!"
        exit 0
    fi
}

function echoStderr(){
    echo "$@" 1>&2
}

function main(){
    urls=(
        health 
        aws/profiles 
        aws/iam/users 
        aws/iam/account 
        aws/iam/alias
        )

    server="http://localhost:3000"
    for url in "${urls[@]}"
    do
        check_status "$server/$url"
    done
    reportResults
}

main