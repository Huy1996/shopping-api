DB_NAME = shopping_cart
NETWORK_NAME = shopping-network
CONTAINER_NAME = shopping-db
SCHEMA_PATH = src/db/migration
DB_SOURCE = "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable"

network:
	docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 || \
        docker network create --driver bridge $(NETWORK_NAME)

postgres:
	docker run -d --name $(CONTAINER_NAME) --network $(NETWORK_NAME) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:14-alpine

create-db:
	docker exec -it shopping-db createdb --username=root --owner=root $(DB_NAME)

drop-db:
	docker exec -it shopping-db dropdb $(DB_NAME)

migrate-up:
	migrate -path $(SCHEMA_PATH) -database $(DB_SOURCE) -verbose up

migrate-up-1:
	migrate -path $(SCHEMA_PATH) -database $(DB_SOURCE) -verbose up 1

migrate-down:
	migrate -path $(SCHEMA_PATH) -database $(DB_SOURCE) -verbose down

migrate-down-1:
	migrate -path $(SCHEMA_PATH) -database $(DB_SOURCE) -verbose down 1

migrate-create:
	migrate create -ext sql -dir $(SCHEMA_PATH) -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

.PHONY: postgres create-db drop-db migrate-up migrate-up-1 migrate-down migrate-down-1 sqlc