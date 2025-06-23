#!/usr/bin/env bash

set -e

# Download the latest gocli binary
echo "Downloading gocli..."
curl -L -o gocli https://github.com/2SSK/gocli/releases/latest/download/gocli

# Make it executable
chmod +x gocli

# Move to /usr/local/bin (may require sudo)
sudo mv gocli /usr/local/bin/gocli

echo "gocli installed successfully!"
echo "You can now run 'gocli --help' from anywhere."
