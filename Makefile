build:
	@go build -o bin/rasta cmd/main.go

test:
	@go test -v y./...

run: build
	@./bin/rasta