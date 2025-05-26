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

.PHONY: all code mocks test mockgen

all: code mocks test

code:
	go run generate/generate.go generate/layout.go generate/requiredParams.go --api=generate/listApis.json

FILES=$(shell for file in `pwd`/cloudstack/*Service.go ;do basename $$file .go ; done)
mocks:
	@for f in $(FILES); do \
		$(MOCKGEN) -destination=./cloudstack/$${f}_mock.go -package=cloudstack -copyright_file="header.txt" -source=./cloudstack/$${f}.go ; \
	done

test:
	go test -v github.com/apache/cloudstack-go/v2/test

MOCKGEN := mockgen
mockgen: ## Install mockgen locally via go install.
	go install go.uber.org/mock/mockgen@latest

