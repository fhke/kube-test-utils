GO_COMMAND = go

.PHONY: help
help: ## Display help info
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: test
test: ## Run all tests
	$(GO_COMMAND) test -v ./...

.PHONY: clean_test_cache
clean_test_cache: ## Clean the go test cache, to force all tests to re-run
	$(GO_COMMAND) clean -testcache

.PHONY: tidy
tidy: ## Tidy go module files
	$(GO_COMMAND) mod tidy
