GRPC_GOOGLE_APIS:=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
SRC:=$(GOPATH)/src
PROTO_PATH:=./proto/*.proto
DB_URI:=mysql://$(ECOMM_CLOUD_DB_CONNECTION_STRING)


default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

install: ## build and install go application executable
	go install -v ./...

clean:  ## go clean
	go clean

build: ## create binary
	go build main.go

run: ## run main.go
	make compile-protobuf
	make check
	go run main.go

test: ## simple running of 'go test' on all directories
	go test ./...

check: ## run linter and vetter on all directories except for vendors
	go list ./... | grep -v /vendor/ | xargs -L1 golint
	go vet

tools: ## Fetch and install required tools and third party requirements
	go get -u github.com/grpc-ecosystem/grpc-gateway/...
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/matryer/moq

compile-protobuf: ## Compile protocol buffer files
	protoc -I. -I$(GOPATH)/src \
	-I$(GRPC_GOOGLE_APIS) --go_out=plugins=grpc:. \
	$(PROTO_PATH)
	protoc -I. \
	-I$(SRC) \
	-I$(GRPC_GOOGLE_APIS) \
	--grpc-gateway_out=logtostderr=true:. \
	$(PROTO_PATH)

compile-binary: ## Create Linux ELF binary
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

generate-swagger: ## Generate swagger docs from protobuf files
	protoc -I. \
	-I$(GOPATH)/src \
	-I$(GRPC_GOOGLE_APIS) \
	--swagger_out=logtostderr=true:. \
	$(PROTO_PATH)

ifeq (dog,$(firstword $(MAKECMDGOALS)))
	echo whee
	$(eval echo whee)
endif


ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "migrate"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):dummy;@:)
endif

ifeq (migrate-create,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "migrate-create"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):dummy;@:)
endif

dummy: ## don't touch this

.PHONY: migrate
##migrate: ## touch $(date +%s)_migration.up $(date +%s)_migration.down
migrate: ## performs migrations up -- run 'migrate help for more info'
	migrate -database '$(DB_URI)' -path ./migrations $(RUN_ARGS)

.PHONY: migrate-create
migrate-create: ## creates migrations with one argument for a suffix
	migrate -database '$(DB_URI)' create -dir migrations -ext .sql $(RUN_ARGS)


