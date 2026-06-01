BINARY = git-bump
VERSION = $(shell git describe --tags --abrev=0 2>/dev/null || echo "dev")

.PHONY: build

build-win:
	go build -ldflags "-X main.version=$(shell git bump manual)" -o git-bump.exe .
	@echo "Build complete."