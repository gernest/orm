#!/usr/bin/env bash

set -e

echo "Generating Person ORM..."
go run main.go -pkg ./example -name Person

echo "Running tests..."
echo "" > coverage.txt
for d in $(go list ./... | grep -v vendor); do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done