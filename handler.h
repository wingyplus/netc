#ifndef HANDLER_H
#define HANDLER_H

typedef struct http_request {
  char *method;
  char *path;
} http_request;

typedef void (*handler_func)(http_request *);

#endif // HANDLER_H
