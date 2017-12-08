#!/usr/bin/env bash

set -e

echo ">> Generating templates binaries"
go generate ./gen/...

echo ">> Installing"
go install ./cmd/orm

echo ">> Cleaning previous generations"
rm -r ./example/*_orm.go || true

echo ">> Generating ORMs..."
go generate ./example/...

echo ">> Running tests..."

if [ -z "${TRAVIS}" ]; then
	go test ./... "$@"
else
	echo "" > coverage.txt
	for d in $(go list ./... | grep -v vendor); do
		go test -v -race -coverprofile=profile.out -covermode=atomic $d
		if [ -f profile.out ]; then
			cat profile.out >> coverage.txt
			rm profile.out
		fi
	done
fi
