start:
	docker compose --env-file ./.env up -d

stop:
	docker compose --env-file ./.env down

logs/app:
	docker compose logs -f --no-log-prefix app

migrate:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=./migrations/ -database ${DB_URI} up

migrate/create:
	docker run -v ./migrations:/migrations --network host migrate/migrate create -ext sql -dir ./migrations $(name)

swagger:
	docker run --rm -v ./:/code ghcr.io/swaggo/swag:v1.16.4 init -g ./cmd/api/main.go

swagger-mac:
	docker run --platform linux/amd64 --rm -v ./:/code ghcr.io/swaggo/swag:latest init -g ./cmd/api/main.go
