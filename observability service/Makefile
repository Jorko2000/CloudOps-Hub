build:
	go build -o bin/observability-service ./cmd/server

run:
	go run ./cmd/server/main.go

docker-build:
	docker build -t cloudops/observability-service .

clean:
	rm -rf bin

fmt:
	go fmt ./...

tidy:
	go mod tidy
