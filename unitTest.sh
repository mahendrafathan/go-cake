#!/usr/bin/env bash

TESTS=$(go test -v -covermode=count -coverprofile=count.txt ./...)
echo "$TESTS"
THRESHOLD=0

if echo "$TESTS" | grep -q "FAIL" ; then
  echo ""
  echo "One or more Unit Tests for app have Failed. Build will now fail. Pipeline will also fail..."
  echo ""
  exit 1
else
  echo ""
  echo "All Unit Tests for application have passed!"
  echo ""
  echo "Running Code Coverage..."
  TOTAL=$(go tool cover -func=./count.txt | grep total | awk '{print substr($3, 1, length($3)-1)}')
  echo "Coverage: $TOTAL"
  if [ $TOTAL \< $THRESHOLD ]
  then
    echo "FAILED: Coverage Below $THRESHOLD%..."
    exit 1
  else
    echo "SUCCESS: Unit Test Coverage"
    exit 0
  fi
fi
