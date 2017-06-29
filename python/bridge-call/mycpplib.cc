#include <iostream>
using namespace std;

int sum(int a, int b)
{
    return a+b;
}

extern "C" {
    int ext_sum(int a, int b) {
        return sum(a, b);
    }
}
