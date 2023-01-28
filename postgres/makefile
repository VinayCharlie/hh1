postgres:
	docker run --name postgres12 -p 8000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
creatmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down



sqlc-help:
	docker run --rm -v /c/Users/vinik/Desktop/gp/Bank:/src -w /src kjconroy/sqlc --help
sqlc-init:
	docker run --rm -v /c/Users/vinik/Desktop/gp/Bank:/src -w /src kjconroy/sqlc init
sqlc-gen:
	docker run --rm -v /c/Users/vinik/Desktop/gp/Bank:/src -w /src kjconroy/sqlc generate


test:
	go test -v -cover ./...




# .PHONY:
# 	postgres createdb dropdb



