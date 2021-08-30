engine:
	@go build -tags musl -a -o engine app/main.go

test:
	@go test ./... -v 

run:
	@docker-compose build
	@docker-compose up -d

stop:
	@docker-compose down

migrate-create:
	@migrate create -ext sql -dir internal/postgres/migrations ${name}

migrate-up:
	@migrate -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_ADDRESS)/$(POSTGRES_DATABASE)?sslmode=disable" \
	-path=internal/postgres/migrations up

migrate-down:
	@migrate -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_ADDRESS)/$(POSTGRES_DATABASE)?sslmode=disable" \
	-path=internal/postgres/migrations down