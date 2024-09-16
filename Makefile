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
	bazel run //:gazelle

.PHONY: build
build: ## build go binary
	bazel build //...

.PHNOY: test
test: test-unit

.PHONY: test-unit
test-unit: ## Run unit tests
	@#bazel test --verbose_failures //... --test_tag_filters=-integration,-external
	go test -v --tags=!integration,!external ./...

.PHONY: test-integration
test-integration: ## Run integration tests
	bazel test --verbose_failures //... --test_tag_filters=integration,-external

.PHONY: coverage
coverage: ## generate coverage report
	@go test -json -coverprofile=cover.out ./... >result.json

.PHONY: gen-pb
gen-pb: ## generate protobuf
	buf generate
	protoc-go-inject-tag -input="./entity/domain/*/*/*.pb.go"

.PHONY: gen-swagger
gen-swagger: ## generate swagger
#	@swag init -q -g impl.go -d ./adapter/restaurant/restful,./entity,./pkg \
#  -o ./api/restaurant/restful --instanceName restaurant_restful --parseDependency
#
#	@swag init -q -g impl.go -d ./adapter/order/restful,./entity,./pkg \
#  -o ./api/order/restful --instanceName order_restful --parseDependency
#
#	@swag init -q -g impl.go -d ./adapter/payment/restful,./entity,./pkg \
#  -o ./api/payment/restful --instanceName payment_restful --parseDependency
#
#	@swag init -q -g impl.go -d ./adapter/user/restful,./entity,./pkg \
#  -o ./api/user/restful --instanceName user_restful --parseDependency
#
#	@swag init -q -g impl.go -d ./adapter/logistics/restful,./entity,./pkg \
#  -o ./api/logistics/restful --instanceName logistics_restful --parseDependency
#
#	@swag init -q -g impl.go -d ./adapter/notify/restful,./entity,./pkg \
#  -o ./api/notify/restful --instanceName notify_restful --parseDependency

### testing
.PHONY: test-e2e-restaurant
test-e2e-restaurant: ## test api
	@k6 run --vus=1 --iterations=1 ./tests/k6/restaurant.e2e.test.js

.PHONY: test-api-order
test-api-order: ## test api
	@k6 run --vus=1 --iterations=1 ./tests/k6/order.api.test.js

.PHONY: test-api-user
test-api-user: ## test api user
	@k6 run --vus=1 --iterations=1 ./tests/k6/user.api.test.js

.PHONY: test-stress
test-stress: ## test load
	@k6 run --env SCENARIO=peak_load ./tests/k6/order.api.test.js --out=cloud

.PHONY: test-load
test-load: ## test stress
	@k6 run --env SCENARIO=average_load ./tests/k6/order.api.test.js --out=cloud

## docker
.PHONY: docker-push
docker-push: ## push docker image
	@bazel run //:push --platforms=@rules_go//go/toolchain:linux_amd64 -- --tag=$(VERSION)

## deployments
DEPLOY_TO := prod
HELM_REPO_NAME := blackhorseya

.PHONY: deploy
deploy: deploy-app deploy-storage ## deploy all

.PHONY: deploy-app
deploy-app: ## deploy app

#.PHONY: deploy-restaurant-restful
#deploy-restaurant-restful: ## deploy restaurant
#	@helm upgrade $(DEPLOY_TO)-godine-restaurant-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-restaurant-restful.yaml
#
#.PHONY: deploy-payment-restful
#deploy-payment-restful: ## deploy payment
#	@helm upgrade $(DEPLOY_TO)-godine-payment-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-payment-restful.yaml
#
#.PHONY: deploy-order-restful
#deploy-order-restful: ## deploy order
#	@helm upgrade $(DEPLOY_TO)-godine-order-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-order-restful.yaml
#
#.PHONY: deploy-user-restful
#deploy-user-restful: ## deploy user
#	@helm upgrade $(DEPLOY_TO)-godine-user-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-user-restful.yaml
#
#.PHONY: deploy-logistics-restful
#deploy-logistics-restful: ## deploy logistics
#	@helm upgrade $(DEPLOY_TO)-godine-logistics-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-logistics-restful.yaml
#
#.PHONY: deploy-notify-restful
#deploy-notify-restful: ## deploy notify
#	@helm upgrade $(DEPLOY_TO)-godine-notify-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-notify-restful.yaml

.PHONY: deploy-storage
deploy-storage: deploy-mariadb deploy-mongodb deploy-redis ## deploy storage

.PHONY: deploy-mariadb
deploy-mariadb: ## deploy mariadb
	@helm upgrade $(DEPLOY_TO)-godine-mariadb bitnami/mariadb \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-mariadb.yaml

.PHONY: deploy-mongodb
deploy-mongodb: ## deploy mongodb
	@helm upgrade $(DEPLOY_TO)-godine-mongodb bitnami/mongodb \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-mongodb.yaml

.PHONY: deploy-redis
deploy-redis: ## deploy redis
	@helm upgrade $(DEPLOY_TO)-godine-redis bitnami/redis \
  --install --namespace $(PROJECT_NAME) \
  --history-max 3 \
  --values ./deployments/$(DEPLOY_TO)/godine-redis.yaml
