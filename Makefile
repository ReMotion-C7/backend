-include .env

build:
	go build -o main.go

run:
	go run main.go

migrate:
	cd cmd/seeder && go run main.go

reset:
	cd cmd/reset && go run main.go

restart-migrate:
	cd cmd/reset && go run main.go && cd ../seeder && go run main.go

compose-build:
	docker compose up --build -d

compose-down:
	docker compose down

compose-running:
	docker ps

compose-log:
	docker-compose logs -f backend-service