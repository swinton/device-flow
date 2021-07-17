.DEFAULT_GOAL := build

build:
	go build -o bin/device-flow main.go

run:
	go run main.go
