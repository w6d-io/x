SHELL=/bin/bash -o pipefail

export GO111MODULE        := on
export PATH               := bin:${PATH}
export PWD                := $(shell pwd)


GO_DEPENDENCIES = golang.org/x/tools/cmd/goimports

define make-go-dependency
  # go install is responsible for not re-building when the code hasn't changed
  bin/$(notdir $1): go.mod go.sum Makefile
		GOBIN=$(PWD)/bin/ go install $1
endef
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))


bin/golangci-lint: Makefile
	bash <(curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh) -d -b .bin v1.28.3

fmt:
	go fmt ./...

vet:
	go vet ./...

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


