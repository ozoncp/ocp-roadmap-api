LOCAL_BIN:=$(CURDIR)/bin

DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ## Test project.
	go test -v ./...

run:
	go run cmd/ocp-roadmap-api/main.go

lint:
	golint ./...

start-db:
	docker-compose up

stop-db:
	docker-compose down

# run goose command
# make goose cmd=create
# make goos cmd=up
gc:
	goose -dir=migrations postgres "postgres://root:root@0.0.0.0:5432/postgres?sslmode=disable" $(cmd)

migrate:
	make gc cmd=up

.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-roadmap-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-roadmap-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-roadmap-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-roadmap-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-roadmap-api/ocp-roadmap-api.proto
		mv pkg/ocp-roadmap-api/github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api/* pkg/ocp-roadmap-api/
		rm -rf pkg/ocp-roadmap-api/github.com
		mkdir -p cmd/ocp-roadmap-api
		cd pkg/ocp-roadmap-api && ls go.mod || go mod init github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api && go mod tidy

.PHONY: generate
generate: .vendor-proto .generate

.PHONY: build
.build:
		go build -o bin/ocp-roadmap-api cmd/ocp-roadmap-api/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-roadmap-api
		cp api/ocp-roadmap-api/ocp-roadmap-api.proto vendor.protogen/api/ocp-roadmap-api/ocp-roadmap-api.proto
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init github.com/ozoncp/ocp-roadmap-api
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
