# git-bump
git-bump is a simple extension for git that mimics the usage of `npm version` for generating new git tags.

## Usage
```bash
git bump <major|minor|patch|manual|version|latest>
```

If no tags are found, `v1.0.0` is created automatically.

## Installation

**Via Go (recommended)**
```bash
go install github.com/neozmmv/git-bump@latest
```

**Windows**
```powershell
irm https://raw.githubusercontent.com/neozmmv/git-bump/master/install.ps1 | iex
```

**Linux**
```bash
curl -fsSL https://raw.githubusercontent.com/neozmmv/git-bump/master/install.sh | bash
```
