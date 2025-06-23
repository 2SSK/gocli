#!/bin/bash

set -e

# Detect OS
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture
if [ "$ARCH" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
    ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# Handle Windows OS (Git Bash or WSL detection)
if [[ "$OS" == "mingw"* || "$OS" == "msys"* || "$OS" == "cygwin"* ]]; then
    OS="windows"
    BINARY="gocli-${OS}-${ARCH}.exe"
    DEST="gocli.exe"
else
    BINARY="gocli-${OS}-${ARCH}"
    DEST="gocli"
fi

# Download binary
URL="https://github.com/2SSK/gocli/releases/latest/download/${BINARY}"

echo "Downloading ${BINARY} from ${URL}..."
curl -L -o ${DEST} ${URL}

# Install
if [[ "$OS" == "windows" ]]; then
    echo "On Windows, please manually move ${DEST} to a directory in your PATH or use it directly."
    echo "You can also add it to PATH using: set PATH=%PATH%;C:\\path\\to\\binary"
else
    chmod +x ${DEST}
    sudo mv ${DEST} /usr/local/bin/gocli
    echo "gocli installed successfully!"
    echo "Run 'gocli --help' to get started."
fi
