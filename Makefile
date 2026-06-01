# Makefile for git-bump
# git bash
# tag before new release for correct versioning
BINARY  = git-bump
VERSION = $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")

.PHONY: build-win build-linux build

build-win:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY).exe .
	@echo "Build complete."

build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY)-linux-arm64 .
	@echo "Build complete."

build:
	@echo "Building for all platforms..."
	$(MAKE) build-win
	$(MAKE) build-linux
	@echo "All builds complete."