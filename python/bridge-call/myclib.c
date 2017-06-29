/***gcc -o myclib.so -shared -fPIC pycall.c*/
#include <stdio.h>
#include <stdlib.h>


int sum(int a, int b)
{
    printf("you input %d and %d\n", a, b);
    return a+b;
}