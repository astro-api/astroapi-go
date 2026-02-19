.PHONY: build test test-cover lint fmt vet tidy generate clean

build:
	go build ./...

test:
	go test -race -count=1 -timeout 120s ./...

test-cover:
	go test -race -count=1 -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

lint:
	golangci-lint run ./...

fmt:
	gofmt -w -s .

vet:
	go vet ./...

tidy:
	go mod tidy && go mod verify

generate:
	go generate ./...

clean:
	rm -f coverage.out coverage.html
