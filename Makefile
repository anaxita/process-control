.PHONY build:
build:
	go build -v ./cmd/killer1c77

.PHONY run:
run:
	go run ./cmd/killer1c77

.DEFAULT_GOAL := run