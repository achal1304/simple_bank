DB_URL= postgresql://achal123:pa55word@localhost:5433/simple_bank?sslmode=disable

postgres:
	docker run --name postgres12 --network bank-network -p 5433:5432 -e POSTGRES_USER=achal123 -e POSTGRES_PASSWORD=pa55word -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=achal123 --owner=achal123 simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

simple_bank:
	docker exec -it postgres12 psql -U achal123 simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" down 1

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

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migratedown migrateup sqlc server mock testcoverage migratedown1 migrateup1 proto evans