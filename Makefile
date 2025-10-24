PHONY: compile run

-include .env

compile:
	go build -o main.go

run:
	go run main.go

migrate:
	go run cmd/seeder/main.go

reset:
	go run cmd/reset/main.go