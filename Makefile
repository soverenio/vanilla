GOTEST ?= go test

export TEST_COUNT ?= 1
export TEST_ARGS ?=
export TEST_DIRS ?= $(shell go list ./... | grep -v /aeshash) # FIXME: aeshash asm signature is broken

.PHONY: all
all: lint test

.PHONY: lint
lint: ## CI lint
	golangci-lint run

.PHONY: test
test:
	$(GOTEST) $(TEST_DIRS) -count $(TEST_COUNT) -race $(TEST_ARGS)
