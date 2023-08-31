postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=achal123 -e POSTGRES_PASSWORD=pa55word -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=achal123 --owner=achal123 simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

simple_bank:
	docker exec -it postgres12 psql -U achal123 simple_bank

migrateup:
	migrate -path db/migration -database postgresql://achal123:pa55word@localhost:5433/simple_bank?sslmode=disable up

migratedown:
	migrate -path db/migration -database postgresql://achal123:pa55word@localhost:5433/simple_bank?sslmode=disable down

sqlc:
	sqlc generate

sqlcdocker:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

testcoverage:
	go test -v -cover -coverprofile=testcoverage.out ./...
	go tool cover -html=testcoverage.out

server:
	go run main.go

mock:
	mockgen -destination=db/mock/store.go -package=mockdb github.com/achal1304/simple_bank/db/sqlc Store	

.PHONY: postgres createdb dropdb migratedown migrateup sqlc server mock testcoverage