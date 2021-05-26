#include "hamming.h"
#include <string.h>

long long compute(const char *lhs, const char *rhs)
{
    int length = 0;
    while(1)
    {
        if (lhs[length] == '\0' && rhs[length] == '\0')
        {
            break;
        }
        if (lhs[length] == '\0' || rhs[length] == '\0')
        {
            return -1;
        }
        length++;
    }

    long long distance = 0;
    for (int i = 0; i < length; i++)
    {
        if (lhs[i] != rhs[i])
        {
            distance++;
        }
    }
    return distance;
}
