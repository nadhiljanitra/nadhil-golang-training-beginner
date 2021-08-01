engine:
	@go build -tags musl -a -o engine app/main.go

test:
	@go test ./... -v 

run:
	@docker-compose build
	@docker-compose up -d

stop:
	@docker-compose down