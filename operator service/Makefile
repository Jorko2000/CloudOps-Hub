APP_NAME=operator-service

build:
	go build -o bin/$(APP_NAME) ./cmd

run:
	go run ./cmd/main.go

docker-build:
	docker build -t cloudops/operator-service:latest .

docker-run:
	docker run cloudops/operator-service:latest

install-crd:
	kubectl apply -f config/crd/application-crd.yaml

install-rbac:
	kubectl apply -f config/rbac/role.yaml

deploy:
	kubectl apply -f config/manager/manager.yaml

clean:
	rm -rf bin

fmt:
	go fmt ./...

tidy:
	go mod tidy
