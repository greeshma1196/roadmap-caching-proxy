build:
	go build -o caching-proxy
run:
	./caching-proxy --port 3000 --origin http://localhost:8080/products
test:
	go test ./... -coverprofile cover.out -v
lint:
	golangci-lint run -v