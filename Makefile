export GO15VENDOREXPERIMENT:=1
export CGO_ENABLED:=0
export GOARCH:=amd64

LOCAL_OS:=$(shell uname | tr A-Z a-z)
GOFILES:=$(shell find . -name '*.go' | grep -v -E '(./vendor)')
GOPATH_BIN:=$(shell echo ${GOPATH} | awk 'BEGIN { FS = ":" }; { print $1 }')/bin
GO_PACKAGE := github.com/odacremolbap/kwtest
LDFLAG_VER := -X $(GO_PACKAGE)/model.version=$(VERSION)
LDFLAG_DATE := -X $(GO_PACKAGE)/model.date=$(DATE)
LDFLAG_GIT := -X $(GO_PACKAGE)/model.gitVersion=$(GIT_VERSION)
LDFLAG_STATIC :=-extldflags "-static"
LDFLAGS=$(LDFLAG_VER) $(LDFLAG_DATE) $(LDFLAG_GIT) $(LDFLAG_STATIC)

all: \
	_output/bin/linux/kwtest \
	_output/bin/darwin/kwtest

release: \
	clean \
	check \
  _output/bin/linux/kwtest

check:
	@find . -name vendor -prune -o -name '*.go' -exec gofmt -s -d {} +
	@go vet $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

install: _output/bin/$(LOCAL_OS)/kwtest
	cp $< $(GOPATH_BIN)

_output/bin/%: $(GOFILES)
	mkdir -p $(dir $@)
	GOOS=$(word 1, $(subst /, ,$*)) go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $@ github.com/odacremolbap/kwtest/cmd/


clean:
	rm -rf _output

.PHONY: all check clean release
