#!/usr/bin/env bash
set -euo pipefail

OWNER="2SSK"
REPO="gocli"

# 1) Detect OS and ARCH
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

echo "Detected OS = ${OS}, ARCH = ${ARCH}"

# 2) Normalize ARCH names
case "${ARCH}" in
    x86_64)           ARCH="amd64"  ;;
    aarch64|arm64)    ARCH="arm64"  ;;
    *)
        echo "❌ Unsupported architecture: ${ARCH}" >&2
        exit 1
        ;;
esac

# 3) Handle Windows-like environments
if [[ "${OS}" == mingw* || "${OS}" == msys* || "${OS}" == cygwin* ]]; then
    OS="windows"
    BINARY="gocli-${OS}-${ARCH}.exe"
    DEST="gocli.exe"
else
    BINARY="gocli-${OS}-${ARCH}"
    DEST="gocli"
fi

URL="https://github.com/${OWNER}/${REPO}/releases/latest/download/${BINARY}"

echo "Downloading ${BINARY} from:"
echo "  ${URL}"

# 4) Download & fail on HTTP errors
if ! curl --fail -L -o "${DEST}" "${URL}"; then
    echo "❌ Download failed. Check that a release asset named '${BINARY}' exists." >&2
    exit 1
fi

# 5) Install
if [[ "${OS}" == "windows" ]]; then
    echo "⚠️  Windows detected — please move '${DEST}' into a PATH directory yourself."
    echo "   For example: move to 'C:\\Windows\\System32' or add to your PATH."
else
    chmod +x "${DEST}"
    echo "Moving to /usr/local/bin/gocli (you may be prompted for your password)..."
    sudo mv "${DEST}" /usr/local/bin/gocli
    echo "✅ gocli installed successfully!"
    echo "Run 'gocli --help' to get started."
fi
