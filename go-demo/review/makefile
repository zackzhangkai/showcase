swag:
	swag init
migrate:
	go run cmd/migrate.go
generate:
	go run cmd/generate.go
hz:
	 hz update -I idl -idl idl/reviews/review.proto -mod reviews

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
		-o reviews -ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildTime=${BUILD_TIME}"
