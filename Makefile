.PHONY: postgres createdb dropdb migrateup migratedown
.SILENT:

build:
	go build -o ./.bin/bot cmd/telegram/main.go

run: build
	./.bin/bot

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root workout_bot

dropdb:
	docker exec -it postgres16 dropdb workout_bot

migrateup:
	migrate -path ./pkg/storage/postgres/migration -database "postgresql://root:secret@localhost:5432/workout_bot?sslmode=disable" -verbose up

migratedown:
	migrate -path ./pkg/storage/postges/migration -database "postgresql://root:secret@localhost:5432/workout_bot?sslmode=disable" -verbose down

migratecreate:
	migrate create -ext sql -dir ./pkg/storage/postges/migration -seq init_schema