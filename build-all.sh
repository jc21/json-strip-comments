#!/bin/bash -e

build() {
	echo "Building $1-$2${3:-} ..."
	GOOS="$1" GOARCH="$2" go build -o "./bin/json-strip-comments-$1-$2${3:-}" ./cmd/json-strip-comments/main.go
}

build darwin arm64
build darwin amd64
build linux	amd64
build linux	arm
build linux	arm64
build windows amd64 ".exe"
