lint:
	golangci-lint run

build:
	go build ./...

test:
	go test ./...
