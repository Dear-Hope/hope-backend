HOST=hope:hope-database-pass@localhost:5432

migrate-create-schema:
	@migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-create-service:
	@migrate create -ext sql -dir scripts/migrations/$(schema) -seq $(name)

migrate-up-service:
	@migrate -source file://scripts/migrations/$(schema) -database "postgres://${HOST}/hope_monolith?sslmode=disable&search_path=$(schema)" -verbose up

migrate-down-service:
	@migrate -source file://scripts/migrations/$(schema) -database "postgres://${HOST}/hope_monolith?sslmode=disable&search_path=$(schema)" -verbose down