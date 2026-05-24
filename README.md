# git-bump

git-bump is a simple extension for git that mimics the usage of `npm version` for generating new git tags.

Usage:
```bash
git bump <major|minor|patch>
```

## Installation

Windows
```pwsh
irm https://raw.githubusercontent.com/neozmmv/git-bump/master/install.ps1 | iex
```

Linux
```bash
curl -fsSL https://raw.githubusercontent.com/neozmmv/git-bump/master/install.sh | bash
```