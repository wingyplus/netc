package main

/*

typedef struct http_request {
        char *method;
        char *path;
} http_request;

typedef void (*handler_func)(http_request *);

inline void call_handler(handler_func fn, http_request *req) {
	fn(req);
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
		var req C.struct_http_request
		req.method = C.CString(r.Method)
		req.path = C.CString(r.URL.Path)
		C.call_handler(C.handler_func(handler), &req)
	})
}

//export http_listen_and_serve
func http_listen_and_serve(listen *C.char) {
	http.ListenAndServe(C.GoString(listen), nil)
}

func main() {}
