postgres:
	docker compose -f db-postgresql/docker-compose.postgresql.yml up -d --build

createdb:
	docker exec -it postgresql createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresql dropdb simple_bank

.PHONY: postgres createdb dropdb