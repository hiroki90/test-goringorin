.PHONY: help
help: ## make コマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## install golangci-lint, goimports, air
	go get -u golang.org/x/tools/cmd/goimports github.com/gin-gonic/gin github.com/go-ozzo/ozzo-validation
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	which air || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

.PHONY: dev-up-all
dev-up-all: ## docker-compose up --build -d
	docker-compose up --build -d

.PHONY: dev-down-all
dev-down-all: ## docker-compose down -v --rmi all
	docker-compose down -v --rmi all

.PHONY: dev-logs
dev-logs: ## docker-compose logs -f
	docker-compose logs -f
