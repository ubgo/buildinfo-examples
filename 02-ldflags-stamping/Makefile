.PHONY: build run clean

# Demo values; in real CI these come from `git describe`, `date -u`, etc.
VERSION    ?= 1.2.3
COMMIT     ?= abc123def456
BUILD_TIME ?= 2026-04-26T12:00:00Z
BRANCH     ?= main

LDFLAGS := \
	-X github.com/ubgo/buildinfo.Version=$(VERSION) \
	-X github.com/ubgo/buildinfo.Commit=$(COMMIT) \
	-X github.com/ubgo/buildinfo.BuildTime=$(BUILD_TIME) \
	-X github.com/ubgo/buildinfo.Branch=$(BRANCH)

build:
	go build -ldflags "$(LDFLAGS)" -o app .

run: build
	./app

# Real-world recipe: pull values from git + the local clock.
build-from-git:
	go build -ldflags "\
		-X github.com/ubgo/buildinfo.Version=$$(git describe --tags --always 2>/dev/null || echo dev) \
		-X github.com/ubgo/buildinfo.Commit=$$(git rev-parse HEAD) \
		-X github.com/ubgo/buildinfo.BuildTime=$$(date -u +%Y-%m-%dT%H:%M:%SZ) \
		-X github.com/ubgo/buildinfo.Branch=$$(git rev-parse --abbrev-ref HEAD)" \
		-o app .

clean:
	rm -f app
