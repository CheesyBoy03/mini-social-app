DOCKER_COMPOSE_FILE ?= docker-compose.yml

dev:
	go run cmd/main.go

migrate-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migrate-down:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1
