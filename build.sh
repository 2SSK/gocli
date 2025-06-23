#!/usr/bin/env bash

set -e

echo "Building for Linux x64..."
GOOS=linux GOARCH=amd64 go build -o gocli-linux-amd64

echo "Building for Linux ARM..."
GOOS=linux GOARCH=arm64 go build -o gocli-linux-arm64

echo "Building for macOS Intel..."
GOOS=darwin GOARCH=amd64 go build -o gocli-darwin-amd64

echo "Building for macOS ARM..."
GOOS=darwin GOARCH=arm64 go build -o gocli-darwin-arm64

echo "Building for Windows x64..."
GOOS=windows GOARCH=amd64 go build -o gocli-windows-amd64.exe

echo "Build completed successfully!"
