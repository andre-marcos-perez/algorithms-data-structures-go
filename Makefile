.PHONY: help
help: ## Show this help message
	@grep -h -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: format
format: ## format code
	@go fmt ./...

.PHONY: lint
lint: ## lint code
	@go vet ./...

.PHONY: test
test: ## test code
	@go test ./...
