# YYDS Music Editor Makefile

.PHONY: build test clean run install lint fmt help

# 默认目标
.DEFAULT_GOAL := help

# 项目信息
PROJECT_NAME := yyds
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 构建标志
LDFLAGS := -X main.Version=$(VERSION) -X main.BuildTime="$(BUILD_TIME)" -X main.GitCommit=$(GIT_COMMIT)

# Go 相关
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod

# 构建目录
BUILD_DIR := build
MAIN_PATH := ./cmd/yyds

build: ## 构建项目
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(PROJECT_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(PROJECT_NAME)"

test: ## 运行测试
	@echo "Running tests..."
	$(GOTEST) -v ./...

test-coverage: ## 运行测试并生成覆盖率报告
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## 清理构建文件
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

run: build ## 构建并运行
	@echo "Running $(PROJECT_NAME)..."
	./$(BUILD_DIR)/$(PROJECT_NAME)

install: ## 安装到 GOPATH/bin
	@echo "Installing $(PROJECT_NAME)..."
	$(GOCMD) install -ldflags "$(LDFLAGS)" $(MAIN_PATH)

lint: ## 运行代码检查
	@echo "Running linter..."
	golangci-lint run

fmt: ## 格式化代码
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

mod-tidy: ## 整理依赖
	@echo "Tidying modules..."
	$(GOMOD) tidy

mod-vendor: ## 创建vendor目录
	@echo "Vendoring dependencies..."
	$(GOMOD) vendor

deps: ## 下载依赖
	@echo "Downloading dependencies..."
	$(GOMOD) download

generate: ## 生成ANTLR代码
	@echo "Generating ANTLR code..."
	@if command -v antlr4 >/dev/null 2>&1; then \
		cd internal/infrastructure/antlr && antlr4 -Dlanguage=Go -o . Music.g4; \
	else \
		echo "Warning: antlr4 not found, skipping code generation"; \
	fi

dev-setup: deps ## 开发环境设置
	@echo "Setting up development environment..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2; \
	fi

release: clean test build ## 发布构建
	@echo "Creating release build..."
	@mkdir -p dist
	@cp $(BUILD_DIR)/$(PROJECT_NAME) dist/
	@cp README.md dist/
	@cp -r examples dist/
	@echo "Release files ready in dist/"

docker-build: ## 构建Docker镜像
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME):$(VERSION) .

help: ## 显示帮助信息
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# 版本信息
version: ## 显示版本信息
	@echo "Project: $(PROJECT_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"