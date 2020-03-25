
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t api-gateway:v0.0.1


.PHONY: run
run:
	go mod download
	CONSUL_ADDRESS=127.0.0.1:8500 MICRO_REGISTRY=consul MICRO_REGISTRY_ADDRESS=127.0.0.1:8500 MICRO_API_HANDLER=http go run main.go

