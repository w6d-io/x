SHELL=/bin/bash -o pipefail

export GO111MODULE  := on
export PATH         := bin:${PATH}
export PWD          := $(shell pwd)
export NEXT_TAG     ?=

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq (,$(shell go env GOOS))
GOOS=$(shell echo $OS)
else
GOOS=$(shell go env GOOS)
endif

ifeq (,$(shell go env GOARCH))
GOARCH=$(shell echo uname -p)
else
GOARCH=$(shell go env GOARCH)
endif


# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef


define install
@[ -f $(1) ] || { \
set -e;\
TMP_DIR=$$(mktemp -d);\
cd $$TMP_DIR ;\
wget -q $(3);\
unzip *.zip $(2);\
mv $(2) $(1);\
rm -rf $$TMP_DIR ;\
}
endef

PROTOC = $(shell pwd)/bin/protoc
ifeq (darwin,$(GOOS))
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-osx-x86_64.zip
else
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip
endif
bin/protoc: ## install protoc locally if necessary.
	$(call install,$(PROTOC),"bin/protoc",$(PROTOC_ZIP))

#.PHONY: bin/golang-ci-lint
bin/golangci-lint:
	bash <(curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh) -d -b .bin v1.28.3

#.PHONY: bin/protoc-gen-go
bin/protoc-gen-go:
	$(call go-get-tool,bin/protoc-gen-go,google.golang.org/protobuf/cmd/protoc-gen-go)


GOIMPORTS = $(shell pwd)/bin/goimports
bin/goimports: ## Download goimports locally if necessary
	$(call go-get-tool,$(GOIMPORTS),golang.org/x/tools/cmd/goimports)


.PHONY: protobuf
protobuf: bin/protoc-gen-go bin/protoc

.PHONY: proto
proto: httpx/test1.pb_test.go

.PHONY: clean
clean:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

httpx/test1.pb_test.go: bin/protoc bin/protoc-gen-go
	$(PROTOC) --go_out=. --go_opt=module=github.com/w6d-io/x httpx/testdata/test1.proto
	mv httpx/test1.pb.go httpx/test1.pb_test.go


.PHONY: format
format: bin/goimports
	goimports -w -local github.com/w6d-io .

.PHONY: lint
lint: bin/golangci-lint
	golangci-lint run -v ./...

.PHONY: test
test: fmt vet
	go test -v -coverpkg=./... -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md --next-tag $(NEXT_TAG)
