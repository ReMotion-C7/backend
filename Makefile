PHONY: compile run migrate reset restart-migrate

-include .env

compile:
	go build -o main.go

run:
	go run main.go

migrate:
	cd cmd/seeder && go run main.go

reset:
	cd cmd/reset && go run main.go

restart-migrate:
	cd cmd/reset && go run main.go && cd ../seeder && go run main.go