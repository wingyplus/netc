package main

/*
#include "handler.h"

void call_handler(handler_func fn, http_request *req) {
	fn(req);
}
*/
import "C"
import "net/http"

func callHandler(handler C.handler_func, r *http.Request) {
	var req C.struct_http_request
	req.method = C.CString(r.Method)
	req.path = C.CString(r.URL.Path)
	C.call_handler(C.handler_func(handler), &req)
}
