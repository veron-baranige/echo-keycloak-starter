MIGRATION_DIR=./database/migrations
DB_CONN_URL=mysql://root:root@tcp\(127.0.0.1:3306\)/business
BINARY_NAME=gocloak

generate: 
	sqlc generate

migrate-init:
	migrate -database ${DB_CONN_URL} -verbose create -ext sql -dir ${MIGRATION_DIR} -seq ${name}

migrate-up:
	migrate -path=${MIGRATION_DIR} -database ${DB_CONN_URL} -verbose up

migrate-down:
	migrate -path=${MIGRATION_DIR} -database ${DB_CONN_URL} -verbose down

migrate-fix:
	migrate -path=${MIGRATION_DIR} -database ${DB_CONN_URL} force ${v}

server:
	go run ./cmd/server/main.go

build:
	@if [ -d bin ]; then rm -rf bin; fi
	go clean
	CGO_ENABLED=0 go build -o bin/${BINARY_NAME} ./cmd/server