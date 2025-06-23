#!/usr/bin/env bash
set -euo pipefail

VERSION=$1

if [[ -z "$VERSION" ]]; then
    echo "❌ Please provide a version, e.g. ./release.sh v1.0.2"
    exit 1
fi

# Build binaries
./build.sh

# Create Git tag
git tag "$VERSION"
git push origin "$VERSION"

# Create GitHub release
gh release create "$VERSION" \
    build/gocli-linux-amd64 \
    build/gocli-linux-arm64 \
    build/gocli-darwin-amd64 \
    build/gocli-darwin-arm64 \
    build/gocli-windows-amd64.exe \
    --title "$VERSION" \
    --notes "Release $VERSION with updated embedded scripts."

echo "✅ Release $VERSION published successfully!"
