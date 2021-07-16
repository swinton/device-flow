.DEFAULT_GOAL := build

build:
	go build -o bin/device-flow device-flow.go

run:
	go run device-flow.go
