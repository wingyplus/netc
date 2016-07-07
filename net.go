package main

// #include "handler.h"
import "C"
import (
	"net/http"
	"unsafe"
)

//export http_handle_func
func http_handle_func(path *C.char, handler unsafe.Pointer) {
	h := C.handler_func(handler)

	http.HandleFunc(C.GoString(path), func(w http.ResponseWriter, r *http.Request) {
		callHandler(h, r)
	})
}

//export http_listen_and_serve
func http_listen_and_serve(listen *C.char) {
	http.ListenAndServe(C.GoString(listen), nil)
}

func main() {}
