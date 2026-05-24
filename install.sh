#!/bin/bash

ARCH=$(uname -m)
OS=$(uname -s)
BASE_URL="https://github.com/neozmmv/git-bump/releases/latest/download"

if [ "$OS" != "Linux" ]; then
    echo "Unsupported OS: $OS. Download manually from https://github.com/neozmmv/git-bump/releases"
    exit 1
fi

if [ "$ARCH" = "x86_64" ]; then
    BINARY="git-bump-linux-amd64"
elif [ "$ARCH" = "aarch64" ]; then
    BINARY="git-bump-linux-arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Downloading git-bump..."
curl -L "$BASE_URL/$BINARY" -o git-bump
chmod +x git-bump
sudo mv git-bump /usr/local/bin/git-bump

echo "Installed! Run: git bump patch"