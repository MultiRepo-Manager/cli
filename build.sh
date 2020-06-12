#!/bin/bash

# Clean previous
rm workspace-*

# Build
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-linux-amd64 .
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-darwin-amd64 .

# Compress
upx -9 workspace-linux-amd64 -o workspace-linux-amd64-upx
upx -9 workspace-darwin-amd64 -o workspace-darwin-amd64-upx

# Install
cp workspace-$(go env GOOS)-$(go env GOARCH)-upx $GOBIN/workspace
