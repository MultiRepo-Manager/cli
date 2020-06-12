#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-linux-amd64 . && upx -9 workspace-linux-amd64
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -v -o workspace-darwin-amd64 . && upx -9 workspace-darwin-amd64
