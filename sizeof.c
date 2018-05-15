/*
 * Print size of standart C types to stdout
 */

#include <stdio.h>  // printf()
#include <stdint.h> // int32_t, int64_t
#include <time.h>   // time_t

int main()
{
#ifndef __cplusplus
  printf("sizeof(void)        = %i\n", (int) sizeof(void));
#endif // __cplusplus

  printf("sizeof(char)        = %i\n", (int) sizeof(char));
  printf("sizeof(short)       = %i\n", (int) sizeof(short));
  printf("sizeof(int)         = %i\n", (int) sizeof(int));
  printf("sizeof(long)        = %i\n", (int) sizeof(long));
  printf("sizeof(long long)   = %i\n", (int) sizeof(long long));
  printf("sizeof(int32_t)     = %i\n", (int) sizeof(int32_t));
  printf("sizeof(int64_t)     = %i\n", (int) sizeof(int64_t));
  printf("sizeof(float)       = %i\n", (int) sizeof(float));
  printf("sizeof(float[2])    = %i\n", (int) sizeof(float[2]));
  printf("sizeof(double)      = %i\n", (int) sizeof(double));
  printf("sizeof(long double) = %i\n", (int) sizeof(long double));
  printf("sizeof(int*)        = %i\n", (int) sizeof(int*));
  printf("sizeof(void*)       = %i\n", (int) sizeof(void*));
  printf("sizeof(time_t)      = %i\n", (int) sizeof(time_t));
  printf("sizeof(size_t)      = %i\n", (int) sizeof(size_t));

  return 0;
}

/*** end of "sizeof.c" file ***/

