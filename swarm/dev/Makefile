.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/etherium/go-etherium/cmd/swarm
	go build -o bin/geth     github.com/etherium/go-etherium/cmd/geth
	go build -o bin/bootnode github.com/etherium/go-etherium/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/etherium/go-etherium/swarm/...
