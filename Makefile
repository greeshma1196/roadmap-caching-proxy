build:
	go build -o caching-proxy
run:
	./caching-proxy
test:
	go test ./... -coverprofile cover.out -v
lint:
	golangci-lint run -v