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
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5432/gophkeeper?sslmode=disable' up 5

.PHONY: migrate-down
migrate-down:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5432/gophkeeper?sslmode=disable' down 5

.PHONY: migrate-up-dev
migrate-up:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5433/gophkeeper_dev?sslmode=disable' up 5

.PHONY: migrate-down-dev
migrate-down:
	migrate -source file://./migrations -database 'postgres://user:password@localhost:5433/gophkeeper_dev?sslmode=disable' down 5