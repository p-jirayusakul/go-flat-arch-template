include .env
export

sqlc:
	sqlc generate

server:
	go run main.go

swag:
	swag init

test:
	go test ./test

mock:
	mockgen -package mockup -destination test/mockup/store.go github.com/p-jirayusakul/go-flat-arch-template/database/sqlc Store

.PHONY: sqlc server swag test mock