#-include ./internal/pkg/config/config.yaml
#export

POSTGRES_USER := postgres
POSTGRES_HOST := localhost
POSTGRES_PORT := 5432
POSTGRES_PASSWORD := root
POSTGRES_SSL_MODE := disable
POSTGRES_DATABASE := arch_db

.PHONY: docker-build
docker-build:
	sudo docker compose up

.PHONY: docker-system-clean
docker-clean:
	sudo docker system prune -af

.PHONY: tidy
tidy:
	go mod tidy && go mod download

.PHONY: run
run: tidy swag-gen
	go run internal/app/main.go

.PHONY: build
build:
	go build -o main internal/app/main.go && ./main

.PHONY: swag-gen
swag-gen:
	swag init -g internal/router/router.go -o internal/docs

.PHONY: create-migration
create-migration:
	migrate create -ext sql -dir internal/migrations -seq "$(name)"

.PHONY: migrate-up
migrate-up:
	migrate -source file://internal/migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${POSTGRES_SSL_MODE} up

.PHONY: migrate-down
migrate-down:
	migrate -source file://internal/migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${POSTGRES_SSL_MODE} down

.PHONY: migration-version
migration-version:
	migrate -database file://internal/migrations - database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${POSTGRES_SSL_MODE} -path migrations version

.PHONY: migrate-dirty
migrate-dirty:
	migrate -path ./internal/migrations/ -database file://internal/migrations - database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${POSTGRES_SSL_MODE} force "$(number)"
