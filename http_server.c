#include "libnet.h"
#include <stdio.h>

void index_handler() {
  printf("hello world\n");
}

int main(int argc, char *argv[]) {
  http_handle_func("/", index_handler);
  http_listen_and_serve(":8080");
}
