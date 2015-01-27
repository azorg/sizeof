/*
 * Print size of standart C types to stdout
 */

#include <stdio.h>  // printf()

#if 0 // FIXME Linux
#  include <stdlib.h> // int32_t, int64_t
#else // FIXME MinGW, Visual C, Linux
#  include <stdint.h> // int32_t, int64_t
#endif

int main()
{
  printf("sizeof(void)      = %i\n", (int) sizeof(void));
  printf("sizeof(char)      = %i\n", (int) sizeof(char));
  printf("sizeof(short)     = %i\n", (int) sizeof(short));
  printf("sizeof(int)       = %i\n", (int) sizeof(int));
  printf("sizeof(long)      = %i\n", (int) sizeof(long));
  printf("sizeof(long long) = %i\n", (int) sizeof(long long));
  printf("sizeof(int32_t)   = %i\n", (int) sizeof(int32_t));
  printf("sizeof(int64_t)   = %i\n", (int) sizeof(int64_t));
  printf("sizeof(float)     = %i\n", (int) sizeof(float));
  printf("sizeof(double)    = %i\n", (int) sizeof(double));
  printf("sizeof(int*)      = %i\n", (int) sizeof(int*));
  printf("sizeof(void*)     = %i\n", (int) sizeof(void*));

  return 0;
}

/*** end of "sizeof.c" file ***/
