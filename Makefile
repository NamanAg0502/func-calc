build:
	@go build -o .bin/func-calc cmd/main.go

run: build
	@.bin/func-calc

test:
	@go test -v ./...

clean:
	@rm -rf .bin
