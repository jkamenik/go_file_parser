export GOPATH=$(shell pwd)/lib

build: get

get:
	go get github.com/lib/pq
