.PHONY: dep
dep: ## Download dependencies
	go mod tidy
	go mod vendor

.PHONY: fmt
fmt: ## Reformat the code
	go fmt ./...

.PHONY: test
test: ## Run the tests
	go test ./... -failfast -coverpkg=./... -coverprofile .testCoverage.txt -race
	@go tool cover -func .testCoverage.txt | grep total | awk '{print "Total coverage: " $$3}'

.PHONY: lint
lint: ## Lint the code
	@command -v golangci-lint >> /dev/null || bash -c "echo 'Installing golangci-lint tool...' && pushd . && cd && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1 && popd"
	golangci-lint run --timeout=3m

.PHONY: help
help: ## Show available commands
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
