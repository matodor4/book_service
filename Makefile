PACKAGE = book_service
APP_NAME ?= $(PACKAGE)
TMPDIR ?= $(shell dirname $$(mktemp -u))
COVER_FILE ?= $(TMPDIR)/$(PACKAGE)-coverage.out

.PHONY: run
run: build ## Start the project
	APP_NAME="$(APP_NAME)" \
	./$(PACKAGE)

.PHONY: build
build: ## Build the project binary
	go build  -o $(APP_NAME) ./cmd/$(PACKAGE)/

.PHONY: test
test: ## Run unit (short) tests
	go test -short -race ./... -coverprofile=$(COVER_FILE)
	go tool cover -func=$(COVER_FILE) | grep ^total

.PHONY: docs
docs: ## Generate swagger docs
	swag init -g cmd/${PACKAGE}/main.go

.PHONY: lint
lint:
	$(info $(M) running linters...)
	golangci-lint --version
	golangci-lint run --timeout 5m0s -v ./...

