#include "armstrong_numbers.h"
#include <math.h>

bool is_armstrong_number(int number)
{
    if (number < 0)
    {
        // negative numbers cannot be armstrong numbers
        return false;
    }
    if (number == 0)
    {
        // the sum of the digits of 0 raised to the power of 1 is 0
        return true;
    }
    int size = trunc(log10(number)) + 1;
    int sum = 0;
    for (int n = number; n != 0; n /= 10)
    {
        sum += pow(n % 10, size);
    }
    return number == sum;
}
