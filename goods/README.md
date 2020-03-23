# User Service

This is the User service

Generated with

```
micro new github.com/liuhaogui/go-micro-mall/goods --namespace=micro.svc.goods --alias=user --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: micro.svc.goods.service.user
- Type: service
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-service
```

Build a docker image
```
make docker
```