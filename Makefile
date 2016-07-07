GO_SOURCES := handler.go net.go

.PHONY: build run

build-go:
	go build -buildmode=c-shared -o libnet.dylib $(GO_SOURCES)

build-c: build-go
	clang -L. -lnet -o http-server http_server.c

run: build-c
	./http-server
