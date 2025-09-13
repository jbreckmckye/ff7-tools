all: clean build

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	go build -ldflags="-w -s" -o build/ .

dev:
	go run .

.PHONY: format
format:
	go fmt .

.PHONY: test
test:
	go test ./internal