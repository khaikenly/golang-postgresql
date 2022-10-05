postgres:
	docker compose -f db-postgresql/docker-compose.postgresql.yml up -d --build

createdb:
	docker exec -it postgresql createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresql dropdb simple_bank

migrateup:
	migrate -path simplebank/db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path simplebank/db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate -f simplebank/sqlc.yaml

.PHONY: postgres createdb dropdb migrateup migratedown sqlc