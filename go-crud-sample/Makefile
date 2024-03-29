.PHONY: help
help: ## make コマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## install goimports, golangci-lint, air
	go get -u golang.org/x/tools/cmd/goimports
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

.PHONY: go-build
go-build: ## 環境依存エラーを回避するため GOOS=linux GOARCH=amd64 CGO_ENABLED=0 を指定
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/server ./cmd/go-crud-sample/main.go

.PHONY: lint
lint: ## linted by golangci-lint
	golangci-lint run -c ./configs/.golangci.yml

.PHONY: goimports
fmt-import: ## formated by goimports
	goimports -w ./

.PHONY: gofmt
fmt: ## formatted by default fmt
	gofmt -w ./

.PHONY: air
air: ## air
	air -c ./configs/.air.toml

.PHONY: mysql
mysql: ## mysql
	docker-compose exec server \
    mysql -u root -pserver_root_pwd -h mysql -P 3306 --protocol=tcp
