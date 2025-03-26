DB_DSN := "postgres://postgres:test@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen:
	@if not exist "./internal/web/${ENTITY}" mkdir "./internal/web/${ENTITY}"
	oapi-codegen -config openapi/.openapi -include-tags ${ENTITY} -package ${ENTITY} openapi/openapi.yaml > ./internal/web/${ENTITY}/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number
