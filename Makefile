.Phony: all

SHELL := /bin/bash # Use bash syntax
default: check-quality test build

ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
WORKDIR=$(shell echo "${PWD}")
APPLICATION_YAML=$(shell echo "$(WORKDIR)/application.yml")
APP_EXECUTABLE="out/codezero"

setup:
	go get -u golang.org/x/lint/golint

check-quality: lint vet

lint:
	@echo "Running Lint ...."
	@if [[ `golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } | wc -l | tr -d ' '` -ne 0 ]]; then \
          golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }; \
          exit 2; \
    fi;

vet:
	@echo "Running Vet checks ...."
	@go vet ./...

test:
	GO111MODULE=on go clean -testcache ./... && go test ./...

build:
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE) .

test-coverage:
	GO111MODULE=on go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
