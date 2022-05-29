#!/usr/bin/env bash

# export SKIP_E2E=true && go test --count=1 -coverprofile=coverage.out ./...

echo -e "\n"
cat coverage.out | \
awk 'BEGIN {cov=0; stat=0;} \
	$3!="" { cov+=($3==1?$2:0); stat+=$2; } \
    END {printf("Total coverage: %.2f%% of statements\n", (cov/stat)*100);}'
echo -e "\n"
