build:
	@go build

run: build
	@go run main.go

test: build
	@go test ./...

testv: build
	@go test ./... -v

lint:
	@golangci-lint run
