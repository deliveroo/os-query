# Go related variables.
GOBASE=$(shell pwd)
export GOBIN=$(GOBASE)/bin
export GOPRIVATE=github.com/deliveroo/*
export GOPROXY=https://proxy.golang.org,off

# Ensure that we use vendored binaries before consulting the system.
GOBIN=$(shell pwd)/bin
export PATH := $(GOBIN):$(PATH)

MODULE = $(shell env GO111MODULE=on go list -m)

# allows passing specific tags to go build (for example musl)
ifdef GO_BUILD_TAGS
GO_BUILD_TAGS_ARG += -tags $(GO_BUILD_TAGS)
endif

gocoverstats=$(GOBIN)/gocoverstats
$(gocoverstats):
	GOBIN=$(GOBIN) go install $(GO_BUILD_TAGS_ARG) gitlab.com/fgmarand/gocoverstats@latest

.PHONY: test
test: ## Run tests
	APP_ENV=test go test -race ./...

.PHONY: test-ci
test-ci: $(go-junit-report) $(gocoverstats) ## Run tests and output as junit-xml
	mkdir -p /tmp/artifacts
	mkdir -p /tmp/test-results
	touch /tmp/test-results/go-test.out
	trap "$(GOBIN)/go-junit-report </tmp/test-results/go-test.out > /tmp/test-results/go-test-report.xml" EXIT; \
	APP_ENV=test go test ${GO_TEST_ARGS} $(GO_BUILD_TAGS_ARG) -coverprofile=/tmp/artifacts/coverage.txt -race ./... 2>&1 | tee /tmp/test-results/go-test.out
	$(GOBIN)/gocoverstats -f /tmp/artifacts/coverage.txt > /tmp/artifacts/test_coverage_stats
	go tool cover -html=/tmp/artifacts/coverage.txt -o /tmp/artifacts/coverage.html