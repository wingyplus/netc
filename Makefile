.PHONY: build run

build-go:
	go build -buildmode=c-shared -o libnet.dylib net.go

build-c: build-go
	clang -L. -lnet -o http-server http_server.c

run: build-c
	./http-server
