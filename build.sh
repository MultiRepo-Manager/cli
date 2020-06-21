#!/bin/bash

# Clean previous
rm -rf workspace-* public

if [[ "$1" == "pkger" ]]; then
  cp -r ../ui/dist public
  pkger -include /public -o server
fi

# Build
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-linux-amd64 .
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-darwin-amd64 .

# Compress
upx -9 workspace-linux-amd64 -o workspace-linux-amd64-upx
upx -9 workspace-darwin-amd64 -o workspace-darwin-amd64-upx

# keep an eye on size
ls -lh workspace-*

# Install
cp workspace-$(go env GOOS)-$(go env GOARCH)-upx $GOBIN/workspace
