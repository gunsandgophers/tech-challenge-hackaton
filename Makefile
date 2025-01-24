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
