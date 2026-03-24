.PHONY: build run test bench fmt clean

BIN_DIR := bin
BIN := $(BIN_DIR)/donut

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN) ./cmd/donut/main.go

run:
	go run ./cmd/donut/main.go $(ARGS)

test:
	go test ./...

bench:
	go test ./donut -run '^$$' -bench . -benchmem

fmt:
	gofmt -w ./cmd ./donut

clean:
	rm -rf $(BIN_DIR)
