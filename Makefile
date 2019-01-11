.DEFAULT_GOAL := all

BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}


#  Commands
#-----------------------------------------------
.PHONY: all
all:
	@gex grapi build -v

.PHONY: setup
setup: dep
	gex --build

.PHONY: dep
dep:
ifdef CI
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	@go get github.com/izumin5210/gex/cmd/gex
	dep ensure -v -vendor-only

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: lint
lint:
ifdef CI
	gex reviewdog -reporter=github-pr-review
else
	gex reviewdog -diff="git diff master"
endif

.PHONY: test
test:
	@go test -v ./...
