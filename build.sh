#!/bin/bash

set -e

APP_NAME="videc"
SOURCE="main.go"

echo "Building executables..."

# Linux
GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}-linux-amd64 $SOURCE
GOOS=linux GOARCH=arm64 go build -o ${APP_NAME}-linux-arm64 $SOURCE

# Windows
GOOS=windows GOARCH=amd64 go build -o ${APP_NAME}-windows-amd64.exe $SOURCE
GOOS=windows GOARCH=arm64 go build -o ${APP_NAME}-windows-arm64.exe $SOURCE

# macOS
GOOS=darwin GOARCH=amd64 go build -o ${APP_NAME}-macos-amd64 $SOURCE
GOOS=darwin GOARCH=arm64 go build -o ${APP_NAME}-macos-arm64 $SOURCE

echo "All builds completed successfully!"