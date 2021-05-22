#include "armstrong_numbers.h"
#include <math.h>

bool is_armstrong_number(int number)
{
    int size = trunc(number <= 0 ? 0 : log10(number)) + 1;
    int sum = 0;
    for (int n = number; n > 0; n /= 10)
    {
        sum += pow(n % 10, size);
    }
    return number == sum;
}
