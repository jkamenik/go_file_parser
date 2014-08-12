export GOPATH=$(shell pwd)/lib

build: get
	go build -o file-parser src/*.go

get:
	go get github.com/lib/pq
