
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t micro-api:v0.0.1


.PHONY: run
run:
	go mod download
	go run main.go

