#include "stddef.h"
#include "stdint.h"
#include "stdlib.h"
#include "string.h"

char *strndup(const char *str, size_t n) {
  if (!str)
    return 0;
  
  size_t length = strlen(str);
  if (length < n)
    n = length;
  
  char *result = (char*)malloc(n + 1);
  memcpy(result, str, n);
  result[n] = '\0';
  return result;
}
