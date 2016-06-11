package main

/*

typedef void (*handler_func)(void);

inline void call_handler(handler_func fn) {
	fn();
}

*/
import "C"
import (
	"net/http"
	"unsafe"
)

//export http_handle_func
func http_handle_func(path *C.char, handler unsafe.Pointer) {
	http.HandleFunc(C.GoString(path), func(w http.ResponseWriter, r *http.Request) {
		C.call_handler(C.handler_func(handler))
	})
}

//export http_listen_and_serve
func http_listen_and_serve(listen *C.char) {
	http.ListenAndServe(C.GoString(listen), nil)
}

func main() {}
