# Makefile for Your Golang Monorepo Project
PROJECT_NAME := $(shell basename $(CURDIR))

# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

VERSION := $(shell git describe --tags --always)

# Targets
.PHONY: all help version
.PHONY: lint clean

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

version: ## show version
	@echo $(VERSION)

.PHONY: dev
dev: ## run dev server
	docker compose up --build

lint: ## run golangci-lint
	@golangci-lint run ./...

clean: ## clean build directory
	@rm -rf cover.out result.json ./deployments/charts/*.tgz
	@rm -rf $(BUILD_DIR)

.PHONY: gazelle
gazelle: ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build
build: ## build go binary
	@bazel build //...

.PHONY: test
test: ## test go binary
	@bazel test --verbose_failures //...

.PHONY: coverage
coverage: ## generate coverage report
	@go test -json -coverprofile=cover.out ./... >result.json

.PHONY: gen-swagger
gen-swagger: ## generate swagger
	@swag init -q -g impl.go -d ./adapter/restaurant/restful,./entity,./pkg \
  -o ./api/restaurant/restful --instanceName restaurant_restful --parseDependency

	@swag init -q -g impl.go -d ./adapter/order/restful,./entity,./pkg \
  -o ./api/order/restful --instanceName order_restful --parseDependency

	@swag init -q -g impl.go -d ./adapter/user/restful,./entity,./pkg \
  -o ./api/user/restful --instanceName user_restful --parseDependency

	@swag init -q -g impl.go -d ./adapter/logistics/restful,./entity,./pkg \
  -o ./api/logistics/restful --instanceName logistics_restful --parseDependency

	@swag init -q -g impl.go -d ./adapter/notify/restful,./entity,./pkg \
  -o ./api/notify/restful --instanceName notify_restful --parseDependency

## docker
.PHONY: docker-push
docker-push: ## push docker image
	@bazel run //:push --platforms=@rules_go//go/toolchain:linux_amd64 -- --tag=$(VERSION)

## deployments
DEPLOY_TO := prod
HELM_REPO_NAME := blackhorseya

.PHONY: deploy
deploy: deploy-restaurant-rest deploy-order-rest deploy-user-rest deploy-logistics-rest deploy-notify-rest ## deploy all

.PHONY: deploy-restaurant-rest
deploy-restaurant-rest: ## deploy restaurant
	@helm upgrade $(DEPLOY_TO)-godine-restaurant-rest $(HELM_REPO_NAME)/godine \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-restaurant-rest.yaml

.PHONY: deploy-order-rest
deploy-order-rest: ## deploy order
	@helm upgrade $(DEPLOY_TO)-godine-order-rest $(HELM_REPO_NAME)/godine \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-order-rest.yaml

.PHONY: deploy-user-rest
deploy-user-rest: ## deploy user
	@helm upgrade $(DEPLOY_TO)-godine-user-rest $(HELM_REPO_NAME)/godine \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-user-rest.yaml

.PHONY: deploy-logistics-rest
deploy-logistics-rest: ## deploy logistics
	@helm upgrade $(DEPLOY_TO)-godine-logistics-rest $(HELM_REPO_NAME)/godine \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-logistics-rest.yaml

.PHONY: deploy-notify-rest
deploy-notify-rest: ## deploy notify
	@helm upgrade $(DEPLOY_TO)-godine-notify-rest $(HELM_REPO_NAME)/godine \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-notify-rest.yaml

