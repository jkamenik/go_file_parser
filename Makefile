export GOPATH=$(shell pwd)/lib

build: get
	go build -o file-parser src/*.go

get:
	go get github.com/lib/pq

test: build
	./file-parser test_data/test.csv.xz
	# ./file-parser test_data/1k_fps.csv.xz
