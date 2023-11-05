SHELL=/bin/bash -o pipefail

export GO111MODULE  := on
export NEXT_TAG     ?=
export CGO_ENABLED   = 1
export LOG_LEVEL     = 2

export PATH         := $(shell pwd)/bin:${PATH}

# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.27.1

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
GOIMPORTS               ?= $(LOCALBIN)/goimports
GOLINT                  ?= $(LOCALBIN)/golint
GOCYCLO                 ?= $(LOCALBIN)/gocyclo
GOLANGCILINT            ?= $(LOCALBIN)/golangci-lint
ENVTEST                 ?= $(LOCALBIN)/setup-envtest
PROTOC                  ?= $(LOCALBIN)/protoc
PROTOC_GEN_GO           ?= $(LOCALBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC      ?= $(LOCALBIN)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY ?= $(LOCALBIN)/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2    ?= $(LOCALBIN)/protoc-gen-openapiv2
PROTOC_GEN_DOC          ?= $(LOCALBIN)/protoc-gen-doc
PROTOC_GEN_GOTAG        ?= $(LOCALBIN)/protoc-gen-gotag
GITCHGLOG               ?= $(LOCALBIN)/git-chglog

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq (,$(shell go env GOOS))
GOOS       = $(shell echo $OS)
else
GOOS       = $(shell go env GOOS)
endif

ifeq (,$(shell go env GOARCH))
GOARCH     = $(shell echo uname -p)
else
GOARCH     = $(shell go env GOARCH)
endif

ifeq (gsed not found,$(shell which gsed))
SEDBIN=sed
else
SEDBIN=$(shell which gsed)
endif

ifeq (darwin,$(GOOS))
GOTAGS = "-tags=dynamic"
else
GOTAGS =
endif

ifndef (,$(NEXT_TAG))
CHGLOG_FLAG = "--next-tag=$(NEXT_TAG)"
else
CHGLOG_FLAG =
endif

##@ Development

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet $(GOTAGS) ./...

.PHONY: test
test: fmt vet envtest
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" go test $(GOTAGS) -v -coverpkg=./... -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total

# Formats the code
.PHONY: format
format: goimports
	$(GOIMPORTS) -w -local github.com/w6d-io,gitlab.w6d.io/w6d cmd internal pkg

.PHONY: lint
lint: golangci-lint golint
	golint ./...
	golangci-lint run -v ./...

.PHONY: cyclo
cyclo: gocyclo
	gocyclo -over 15 .

## Tool Versions
PROTOC_TOOLS_VERSION     ?= 3.20.3

PROTOC_ARCH ?= x86_64
PROTOC_OS   ?= linux

PROTOC = $(shell pwd)/bin/protoc
ifeq (darwin,$(GOOS))
PROTOC_OS = osx
endif

PROTOC_ZIP ?= https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_TOOLS_VERSION)/protoc-$(PROTOC_TOOLS_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip

.PHONY: protobuf
protobuf: protoc protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-gateway protoc-gen-openapiv2 protoc-gen-doc protoc-gen-gotag

.PHONY: protoc
protoc: $(PROTOC)
$(PROTOC): $(LOCALBIN) ## install protoc locally if necessary.
	@test -s $(LOCALBIN)/protoc || $(call install,$(PROTOC),bin/protoc,$(PROTOC_ZIP))

.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-go || GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-go-grpc || GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: protoc-gen-grpc-gateway
protoc-gen-grpc-gateway: $(PROTOC_GEN_GRPC_GATEWAY)
$(PROTOC_GEN_GRPC_GATEWAY): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-grpc-gateway || GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

.PHONY: protoc-gen-openapiv2
protoc-gen-openapiv2: $(PROTOC_GEN_OPENAPIV2)
$(PROTOC_GEN_OPENAPIV2): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-openapiv2 || GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

.PHONY: protoc-gen-doc
protoc-gen-doc: $(PROTOC_GEN_DOC)
$(PROTOC_GEN_DOC): $(LOCALBIN)
	@go mod download github.com/pseudomuto/protoc-gen-doc
	@test -s $(LOCALBIN)/protoc-gen-doc || GOBIN=$(LOCALBIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

.PHONY: protoc-gen-gotag
protoc-gen-gotag: $(PROTOC_GEN_GOTAG)
$(PROTOC_GEN_GOTAG): $(LOCALBIN)
	@test -s $(LOCALBIN)/protoc-gen-gotag || GOBIN=$(LOCALBIN) go install github.com/srikrsna/protoc-gen-gotag@latest

.PHONY: chglog
chglog: $(GITCHGLOG) ## Download git-chglog locally if necessary
$(GITCHGLOG): $(LOCALBIN)
	@test -s $(LOCALBIN)/git-chglog || GOBIN=$(LOCALBIN) go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	@test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: goimports
goimports: $(GOIMPORTS) ## Download goimports locally if necessary
$(GOIMPORTS): $(LOCALBIN)
	@test -s $(LOCALBIN)/goimports || GOBIN=$(LOCALBIN) go install golang.org/x/tools/cmd/goimports@latest

.PHONY: golint
golint: $(GOLINT) ## Download golint locally if necessary
$(GOLINT): $(LOCALBIN)
	@test -s $(LOCALBIN)/golint || GOBIN=$(LOCALBIN) go install golang.org/x/lint/golint@latest

.PHONY: gocyclo
gocyclo: $(GOCYCLO) ## Download gocyclo locally if necessary
$(GOCYCLO): $(LOCALBIN)
	@test -s $(LOCALBIN)/gocyclo || GOBIN=$(LOCALBIN) go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

.PHONY: golangci-lint
golangci-lint: $(GOLANGCILINT) ## Download golangci-lint locally if necessary
$(GOLANGCILINT): $(LOCALBIN)
	@test -s $(LOCALBIN)/golangci-lint || GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@@latest

.PHONY: proto
proto: httpx/test1.pb_test.go mongox/test1.pb_test.go

httpx/test1.pb_test.go: protoc protoc-gen-go
	$(PROTOC) --go_out=. --go_opt=module=github.com/w6d-io/x httpx/testdata/test1.proto
	mv httpx/test1.pb.go httpx/test1.pb_test.go

mongox/test1.pb_test.go: protoc protoc-gen-go
	$(PROTOC) --proto_path=./mongox/testdata --go_out=. --go_opt=module=github.com/w6d-io/x mongox/testdata/test1.proto
	mv mongox/test1.pb.go mongox/test1.pb_test.go


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
