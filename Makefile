postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5433:5432 -d postgres:16-alpine
createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root golang

dropdb:
	docker exec -it postgres16 dropdb golang

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

