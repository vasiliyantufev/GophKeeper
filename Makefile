.PHONY: up
up:
	docker-compose up

.PHONY: dev-up
dev-up:
	docker-compose -f docker-compose.dev.yaml up

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -seq -dir ./migrations name

.PHONY: migrate-up
migrate-up:
	migrate -source file://./migrations -database 'postgresql://localhost:5432/postgresql?user=other&password=password' up 4

.PHONY: migrate-down
migrate-down:
	migrate -source file://./migrations -database 'postgresql://localhost:5432/postgresql?user=other&password=password' down 4