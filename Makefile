include ./.env

start: dc-up app-up run migrate-up

app-build:
	go build -o ./build/app ./cmd/app

run:
	./build/app

app-up: app-build run

dc-up-local:
	docker-compose -f ./docker-compose.yml --env-file .env up -d database cache_redis

dc-up:
	docker-compose -f ./docker-compose.yml up -d

dc-down:
	docker-compose -f ./docker-compose.yml down


migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

migrate-up: migrate-install
	migrate -path ./migration -database="${DB_SCHEME}://${DB_LOGIN}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" up

migrate-new: migrate-install
	migrate create -ext sql -dir ./migration "$(name)"
