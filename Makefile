postgres:
	docker run -d --name postgres96 -p 3356:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nader123 postgres

createdb:
	docker exec -it postgres96 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres96 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgres://root:nader123@localhost:3356/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:nader123@localhost:3356/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test