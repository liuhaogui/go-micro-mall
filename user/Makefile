GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	cd proto/user && protoc --proto_path=.:${GOPATH}/src --go_out=. --micro_out=. user.proto && cd -

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o user-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:v0.0.1

.PHONY: run
run:
	go mod download
	CONSUL_ADDRESS=127.0.0.1:8500 MICRO_REGISTRY=consul MICRO_REGISTRY_ADDRESS=127.0.0.1:8500 go run main.go



