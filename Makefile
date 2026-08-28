.PHONY: build run

run:
	go run cmd/handcode/main.go

build:
	go build -o cmd/handcode/handcode cmd/handcode/main.go 