DB_SOURCE = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
DB_NAME = shopping_cart
NETWORK_NAME = shopping-network
CONTAINER_NAME = shopping-db

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
	migrate -path db/migration -database $(DB_SOURCE) -verbose up

migrate-up-1:
	migrate -path db/migration -database $(DB_SOURCE) -verbose up 1

migrate-down:
	migrate -path db/migration -database $(DB_SOURCE) -verbose down

migrate-down-1:
	migrate -path db/migration -database $(DB_SOURCE) -verbose down 1

migrate-create:
	migrate create -ext sql -dir src/db/migration -seq $(name)

.PHONY: postgres create-db drop-db migrate-up migrate-up-1 migrate-down migrate-down-1