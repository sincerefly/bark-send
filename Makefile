BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
GO_VERSION := $(shell go version)
GIT_BRANCH := $(shell git symbolic-ref --short -q HEAD)
GIT_BRANCH ?= $(shell git describe --abbrev=0 --tag)
GIT_COMMIT_HASH := $(shell git rev-parse --verify HEAD)
GO_FLAGS := -v -ldflags="-w -s -X 'bark-send/common.Build=$(BUILD_DATE)' -X 'bark-send/common.Commit=$(GIT_COMMIT_HASH)' -X 'bark-send/common.Branch=$(GIT_BRANCH)' -X 'bark-send/common.Platform=${GO_VERSION}'"

bl_mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(GO_FLAGS) -v -o bin/bark-send-darwin

bl_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_FLAGS) -v -o bin/bark-send-linux
