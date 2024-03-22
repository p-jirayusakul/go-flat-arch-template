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
	mockgen -package mockup -destination test/mockup/api.go github.com/p-jirayusakul/go-flat-arch-template/external ExternalAPI

.PHONY: sqlc server swag test mock