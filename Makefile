# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: code mock-gen test

code:
	go run generate/generate.go generate/layout.go --api=generate/listApis.json

FILES=$(shell for file in `pwd`/cloudstack/*Service.go ;do basename $$file .go ; done)
mock-gen:
	@for f in $(FILES); do \
		$(MOCKGEN) -destination=./cloudstack/$${f}_mock.go -package=cloudstack -copyright_file="header.txt" -source=./cloudstack/$${f}.go ; \
	done

test:
	go test -v ./...

MOCKGEN := $(shell pwd)/bin/mockgen
mockgen: ## Download conversion-gen locally if necessary.
	$(call go-get-tool,$(MOCKGEN),github.com/golang/mock/mockgen)

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
