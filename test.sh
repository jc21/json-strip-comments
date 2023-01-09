#!/bin/bash -e

# check for tparse
if hash tparse 2>/dev/null; then
	go test -json -cover ./... | tparse
else
	go test -bench=. -cover -v ./...
fi
