include .env

docker-up:
	@docker compose --env-file .env up --build

#	Migrations
MIGRATE_PATH=./pkg/migrate/migrations

migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATE_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATE_PATH) -database=$(DATABASE_URL) up

migrate-down:
	@migrate -path=$(MIGRATE_PATH) -database=$(DATABASE_URL) down ${filter-out $@,${MAKECMDGOALS}}