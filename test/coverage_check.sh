#!/usr/bin/env bash

# echo -e "\n"
# cat coverage.out | \
# awk 'BEGIN {cov=0; stat=0;} \
# 	$3!="" { cov+=($3==1?$2:0); stat+=$2; } \
#     END {printf("Total coverage: %.2f%% of statements\n", (cov/stat)*100);}'
# echo -e "\n"

TESTCOVERAGE_THRESHOLD=91
echo -e "\nQuality Gate: Test Coverage Check"
echo "Threshold             : $TESTCOVERAGE_THRESHOLD %"
totalCoverage=$(go tool cover -func=build/coverage/coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')
echo "Current test coverage : $totalCoverage %"
echo ""
if (( $(echo "$totalCoverage $TESTCOVERAGE_THRESHOLD" | awk '{print ($1 > $2)}') )); then
    echo "✅ - OK"
else
    echo "Current test coverage is below threshold."
    echo "Please add more unit tests"
    echo "❌ - Failed"
    exit 1
fi
echo ""
