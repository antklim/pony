.PHONY: build
build: go-build ## Build Go services

.PHONY: clean
clean: go-clean ## Clean Go build cache and dependencies

.PHONY: deps
deps: go-deps ## Install Go dependencies

.PHONY: test
test: go-test ## Run tests

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build:
	@echo "Building Pony..."
	@rm -rf build
	@mkdir build
	go build -o build -v ./...
	@echo "Pony is ready at ./build"

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	go clean -cache

go-clean-test-cache:
	go clean -testcache

go-clean-deps:
	go mod tidy

go-deps:
	go mod download

go-test:
	go test -v -tags="unit" ./...

pony-clean: ## Clean pony site
	@echo "Pony clean ..."
	@rm -fr _build

pony-build: ## Build pony site
	@$(MAKE) pony-clean
	@echo "Pony build ..."
	@mkdir _build
	@build/pony build -o _build -t example -m example/pony.yaml

pony-site-map: ## Run pony site map
	@build/pony run -t example -m example/pony.yaml --sitemap -a :9000

.DEFAULT_GOAL := help