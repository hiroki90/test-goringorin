.PHONY: help
help: ## make コマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: dev-up-all
dev-up-all: ## docker-compose up --build -d
	docker-compose up --build -d

.PHONY: dev-down-all
dev-down-all: ## docker-compose down -v --rmi all
	docker-compose down -v --rmi all

.PHONY: dev-logs
dev-logs: ## docker-compose logs -f
	docker-compose logs -f
