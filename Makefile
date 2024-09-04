GOTEST ?= go test
GOARCHFORCED ?= # Way to force GOARCH to amd64 for M1 chips

export TEST_COUNT ?= 1
export TEST_ARGS ?=

.PHONY: all
all: lint test

.PHONY: lint
lint: ## CI lint
	GOARCH=$(GOARCHFORCED) golangci-lint run

.PHONY: test
test:
	GOARCH=$(GOARCHFORCED) $(GOTEST) ./... -count $(TEST_COUNT) -race $(TEST_ARGS)

.PHONY: init-devbox
init-devbox:
	go mod tidy
