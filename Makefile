GO_ENV=CGO_ENABLED=0 GO111MODULE=on
GO=$(GO_ENV) $(shell which go)
VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_DATE = `date +%FT%T%z`
BIN=simple-backend
DIR_SRC=./cmd/app
PROJ_PATH=github/kunhou/simple-backend
GO_FLAGS=-ldflags="-X '$(PROJ_PATH)/cmd.Version=$(VERSION)' -X '$(PROJ_PATH)/cmd.GitCommitSha=$(GIT_COMMIT)' -X '$(PROJ_PATH)/cmd.BuildDate=$(BUILD_DATE)'"
WIRE_VERSION=v0.5.0
MOCKGEN_VERSION=v1.6.0

.PHONY: sql-generate sql-rehash run build generate gen-grpc

sql-generate: 
	atlas migrate diff --env gorm m.up
sql-rehash:
	atlas migrate hash --env gorm

run: generate
	@$(GO) run $(GO_FLAGS) $(DIR_SRC)

build: generate
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)

generate: gen-grpc
	@$(GO) get github.com/google/wire/cmd/wire@$(WIRE_VERSION)
	@$(GO) get github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)
	@$(GO) install github.com/google/wire/cmd/wire@$(WIRE_VERSION)
	@$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)
	@$(GO) generate ./...
	@$(GO) mod tidy
	@$(MAKE) gen-grpc

gen-grpc:
	@protoc --go_out=. --go-grpc_out=. deliver/grpc/proto/*.proto