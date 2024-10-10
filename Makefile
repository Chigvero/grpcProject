include local.env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=${MIGRATION_DIR}
LOCAL_MIGRATION_DSN=${PG_DSN}


install-deps:
	make install-proto-deps
	make install-goose
	make install-golangci-lint

generate:
	make generate-note-api

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


##linter
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml
##linter



##grpc
install-proto-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=./bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	api/note_v1/note.proto
##grpc



##Migrations
install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

local-migration-status:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(LOCAL_MIGRATION_DSN) $(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) status -v

local-migration-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(LOCAL_MIGRATION_DSN) $(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) up -v

local-migration-down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(LOCAL_MIGRATION_DSN) $(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) down -v
##Migrations



##CI/CD
##Пример того как сделать все мануально
##Надо Автоматизировать это все через CI/CD
docker-build-and-push:
	docker buildx create --use
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/chigvero/my_registry:v0.0.1 . --load
	##docker login -u token -p CRgAAAAAaNthevgWAUiQ4kYsHGyKMNXR4hgSFICc==  cr.selcloud.ru/chigvero
	docker push cr.selcloud.ru/chigvero/my_registry:v0.0.1
##CI/CD



##Manual build and push
build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	scp service_linux root@109.172.89.231:
##Manual build and push