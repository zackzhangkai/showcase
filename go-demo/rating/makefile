update_gen:
	kitex -I idl/ -type protobuf idl/rating.proto

# Environment variables set when running the Go compiler.
VERSION ?= v1.0.0-SNAPSHOT
COMMIT_HASH ?= $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date "+%Y-%m-%dUTC%H:%M:%S")

GO ?= go
GOOS ?= $(shell ${GO} env GOOS)
GOARCH ?= $(shell ${GO} env GOARCH)
GO_BUILD_ENVVARS = \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	CGO_ENABLED=0 \

## build: build binary file
build:
	@echo Building...
	${GO_BUILD_ENVVARS} ${GO} build -mod vendor \
		-o ratings -ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildTime=${BUILD_TIME}"