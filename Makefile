build: build-go ## Build the app

dep-upgrade: dep-upgrade-go ## Upgrades dependencies

dep-upgrade-go:
	@go get -u
	@go mod tidy

build-go:
	@go build -v -ldflags="-s -w -X 'github.com/michaelcoll/version-badge/cmd.version=v0.0.0'" .

.PHONY: test
test: ## Launch go tests and linter
	@go test -vet=all ./...

.PHONY: coverage
coverage: ## Launch go tests with coverage
	@go test -vet=all -covermode atomic -coverprofile=coverage.out ./...

run-back: ## Launch the backend
	@go run . serve

run-docker:
	docker compose up --build

gen: generate ## Generate all the code

.PHONY: generate
generate:
	@go generate internal/back/domain/callers.go

.PHONY: help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'