.PHONY: build test clean fmt lint

build:
	go build -v ./...

test:
	go test -v ./...

clean:
	go clean
	rm -f coverage.out

fmt:
	go fmt ./...

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
