ROOT_DIR = $(shell pwd)
PATH := $(ROOT_DIR)/bin:$(shell go env GOPATH)/bin:$(PATH)
COLOR := "\e[1;36m%s\e[0m\n"
PROTO_ROOT := $(ROOT_DIR)/api/proto
PROTO_EXAMPLE_ROOT := $(ROOT_DIR)/examples/proto

##### Build tools #####
.PHONY: install-build-tools
install-build-tools:
	@printf $(COLOR) "Installing build tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest

##### Proto to go generation #####
.PHONY: codegen-proto
codegen-proto:
	@printf $(COLOR) "Generating proto..."
	@(cd $(PROTO_ROOT) && buf generate)

##### Example proto to go generation #####
.PHONY: codegen-proto-example-local
codegen-proto-example-local:
	make build-codegenerator
	@printf $(COLOR) "Generating proto..."
	@(cd $(PROTO_EXAMPLE_ROOT) && buf generate --template buf.gen.local.yaml)

##### Build protoc-gen-natsrpcgo #####
.PHONY: build-codegenerator
build-codegenerator:
	@printf $(COLOR) "Building generator plugin..."
	@go build -o bin/ ./cmd/protoc-gen-rpc-errormapper-go

##### Build #####
.PHONY: build
build:
	@printf $(COLOR) "Building..."
	@go build -o ./bin/ -v ./...

##### Lint #####
.PHONY: lint
lint: lint-proto lint-go

.PHONY: lint-go
lint-go:
	@printf $(COLOR) "Lint go..."
	@golangci-lint run ./...

.PHONY: lint-proto
lint-proto:
	@printf $(COLOR) "Lint proto..."
	@(cd $(PROTO_ROOT) && buf lint)

##### Vulnerability and Security #####
.PHONY: check-security
check-security:
	@printf $(COLOR) "Checking security..."
	@gosec ./...

.PHONY: check-vulnerabilities
check-vulnerabilities:
	@printf $(COLOR) "Checking vulnerabilities..."
	@govulncheck ./...

##### Test #####
.PHONY: test
test:
	@printf $(COLOR) "Testing..."
	@go test -race ./...

##### Dependencies #####
.PHONY: install
install:
	@printf $(COLOR) "Install Dependencies..."
	@go mod download
