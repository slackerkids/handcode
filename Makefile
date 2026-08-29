.PHONY: build run

run:
	go run cmd/handcode/main.go

build:
	go build -o build/handcode cmd/handcode/main.go 