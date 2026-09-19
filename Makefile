build:
	@go build -o bin/rasta cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/rasta