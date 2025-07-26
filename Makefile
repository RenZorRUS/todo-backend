# Development tools versions
GOLANGCI_LINT_VERSION ?= 2.3.0
PRE_COMMIT_VERSION    ?= 4.2.0
UPX_VERSION           ?= 5.0.2

# Operation system specific variables
OS := $(shell uname -s 2>/dev/null || echo Windows)

# Paths to local `bin` project directory
DEV_BIN     := $(CURDIR)/bin
DEV_TOOLS_PATHS := $(DEV_BIN)/pre-commit $(DEV_BIN)/golangci-lint $(DEV_BIN)/upx
export PATH := $(DEV_BIN):$(PATH)

# Values for running app in production
BINARY_NAME  := todo-backend
MAIN_DIR     := ./src/cmd/api
BUILD_DIR    := ./bin
BINARY_PATH  := $(BUILD_DIR)/$(BINARY_NAME)

# Testing values
COVERAGE_DIR := ./coverage
COVERAGE_EXCLUDE := $(shell go list ./... | grep -E 'domains|ports|consts|cmd|tests|errs')

# Go build flags
GOFLAGS     := -ldflags="-s -w"
CGO_ENABLED := 0

# Platform-specific installers
ifeq ($(OS),Linux)
	BREW_CMD := /bin/false
	CHOCO_CMD := /bin/false
else ifeq ($(OS),Darwin)
	BREW_CMD := brew
	CHOCO_CMD := /bin/false
else # Windows
	BREW_CMD := /bin/false
	CHOCO_CMD := choco
endif

# Phony targets (action names)
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make dev-tools-install  - ⬇️  installs pre-commit & golangci-lint"
	@echo "  make pre-commit-install - ⚙️  set up pre-commit hooks in your local .git directory"
	@echo "  make clean-dev-tools    - 🚮 removes development tools"
	@echo "  make check              - 🔎 checks if all development tools are installed"
	@echo "  make lint               - 🔥 runs Golang linters according to rules in .golangci.yaml"
	@echo "  make fmt                - 🧹 formats Golang code according to rules in .golangci.yaml"
	@echo "  make tidy               - 🧹 checks, cleans and updates go.mod and go.sum files"
	@echo "  make audit              - 🔥 runs Golang linters and others quality checks"
	@echo "  make unit-test          - 🧪 runs unit tests and creates coverage report in $(COVERAGE_DIR)"
	@echo "  make clean              - 🚮 removes binary executable: $(BINARY_PATH)"
	@echo "  make prod-linux-amd64   - 🚀 builds binary executable: $(BINARY_PATH)-linux-amd64 from: $(MAIN_DIR)"
	@echo "  make build              - 🚀 builds binary executable: $(BINARY_PATH) from: $(MAIN_DIR)"
	@echo "  make run                - 🔥 runs binary executable: $(BINARY_PATH)"

# Pre-commit installation target
$(DEV_BIN)/pre-commit:
	@echo "==> ⬇️ Installing pre-commit@$(PRE_COMMIT_VERSION)"
	@mkdir -p $(DEV_BIN)
ifeq ($(BREW_CMD),brew)
	@$(BREW_CMD) install pre-commit
	@ln -sf $$(which pre-commit) $(DEV_BIN)/pre-commit
else ifeq ($(CHOCO_CMD),choco)
	@powershell -Command "choco install pre-commit --version $(PRE_COMMIT_VERSION) -y --no-progress"
	@ln -sf $$("C:/ProgramData/chocolatey/bin/pre-commit.exe") $(DEV_BIN)/pre-commit
else
# Fallback: pip-less install via pre-commit-standalone (if available)
	@curl -sSL https://github.com/pre-commit/pre-commit/releases/download/v$(PRE_COMMIT_VERSION)/pre-commit-$(PRE_COMMIT_VERSION)-py2.py3-none-any.whl \
		-o $(DEV_BIN)/pre-commit.zip && unzip -q -o $(DEV_BIN)/pre-commit.zip -d $(DEV_BIN)
endif
	@touch $(DEV_BIN)/pre-commit

# Golangci-lint installation target
$(DEV_BIN)/golangci-lint:
	@echo "==> ⬇️ Installing golangci-lint@$(GOLANGCI_LINT_VERSION)"
	@mkdir -p $(DEV_BIN)
ifeq ($(BREW_CMD),brew)
	@$(BREW_CMD) install golangci-lint
	@ln -sf $$(which golangci-lint) $(DEV_BIN)/golangci-lint
else ifeq ($(CHOCO_CMD),choco)
	@powershell -Command "choco install golangci-lint --version $(GOLANGCI_LINT_VERSION) -y --no-progress"
	@ln -sf $$("C:/ProgramData/chocolatey/bin/golangci-lint.exe") $(DEV_BIN)/golangci-lint
else
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(DEV_BIN) v$(GOLANGCI_LINT_VERSION)
endif
	@touch $(DEV_BIN)/golangci-lint

# Upx installation target
$(DEV_BIN)/upx:
	@echo "==> ⬇️ Installing upx"
	@mkdir -p $(DEV_BIN)
ifeq ($(BREW_CMD),brew)
	@$(BREW_CMD) install upx
	@ln -sf $$(which upx) $(DEV_BIN)/upx
else ifeq ($(CHOCO_CMD),choco)
	@powershell -Command "choco install upx --version $(UPX_VERSION) -y --no-progress"
	@ln -sf $$("C:/ProgramData/chocolatey/bin/upx.exe") $(DEV_BIN)/upx
else
	@curl --location --output upx-$(UPX_VERSION)-amd64_linux.tar.xz "https://github.com/upx/upx/releases/download/v$(UPX_VERSION)/upx-$(UPX_VERSION)-amd64_linux.tar.xz" && \
    tar -xJf upx-$(UPX_VERSION)-amd64_linux.tar.xz && \
    cp upx-$(UPX_VERSION)-amd64_linux/upx /bin/
endif
	@touch $(DEV_BIN)/upx

# Development tools installation
.PHONY: dev-tools-install pre-commit-install clean-dev-tools check
dev-tools-install: $(DEV_TOOLS_PATHS)

pre-commit-install: $(DEV_BIN)/pre-commit
	@pre-commit install

clean-dev-tools:
	@rm $(DEV_TOOLS_PATHS)

check: dev-tools-install pre-commit-install lint
	@echo "✅ All development tools are installed!"

# Go code formatters and linters
.PHONY: lint fmt tidy
lint: $(DEV_BIN)/golangci-lint
	@golangci-lint run

fmt: $(DEV_BIN)/golangci-lint
	@golangci-lint fmt

tidy: fmt
	@go mod tidy -v -x

# Quality checks
.PHONY: audit unit-tests
audit: fmt lint tidy
	@go vet ./...
	@go mod verify

unit-tests:
	@mkdir -p $(COVERAGE_DIR)
	@go test -race -coverprofile=$(COVERAGE_DIR)/cover.out \
		$(filter-out $(COVERAGE_EXCLUDE),$(shell go list ./...))
	@go tool cover -html=$(COVERAGE_DIR)/cover.out -o $(COVERAGE_DIR)/cover.html

# Go binary builders
.PHONY: clean prod prod-linux-amd64 build run
clean:
	@rm -f $(BINARY_PATH) $(BINARY_PATH)-linux-amd64

prod: clean
	CGO_ENABLED=$(CGO_ENABLED) \
		go build $(GOFLAGS) -o $(BINARY_PATH) $(MAIN_DIR)
	@upx -9 $(BINARY_PATH)

prod-linux-amd64: clean
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=$(CGO_ENABLED) \
		go build $(GOFLAGS) -o $(BINARY_PATH)-linux-amd64 $(MAIN_DIR)
	@upx -9 $(BINARY_PATH)-linux-amd64

build: tidy
	@go build $(GOFLAGS) -o $(BINARY_PATH) $(MAIN_DIR)

run: build
	@$(BINARY_PATH)
