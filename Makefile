PHONY: compile run

-include .env

compile:
	go build -o main.go

run:
	go run main.go