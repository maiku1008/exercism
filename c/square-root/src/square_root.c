#include "square_root.h"
#include <stdlib.h>

uint16_t square_root(uint16_t n)
{
    uint16_t c = 0, sum = 0;
    while (sum < n)
    {
        c++;
        sum += 2 * c - 1;
    }
    return c;
}
