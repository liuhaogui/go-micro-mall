
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
    
	cd proto/hello && protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. hello.proto && cd -
    

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t hello-srv:v0.0.1


.PHONY: run
run:
	go mod download
	MICRO_REGISTRY=consul MICRO_REGISTRY_ADDRESS=127.0.0.1:8500 go run main.go

