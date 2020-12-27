build:
	go build  -o bin/inventory cmd/main.go

run:
	bin/inventory

all: build run