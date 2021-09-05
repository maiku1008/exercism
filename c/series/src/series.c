#include "series.h"
#include <stdlib.h>
#include <string.h>

slices_t slices(char *input_text, unsigned int substring_length)
{
    slices_t result = {0};

    int len = strlen(input_text);
    if (len == 0 || substring_length == 0)
    {
        return result;
    }
    result.substring_count = 1 + (len - substring_length);
    result.substring = calloc(sizeof(char *), result.substring_count);

    for (int i = 0; i <= (int)(len - substring_length); i++)
    {
        result.substring[i] = calloc(sizeof(char), substring_length);
        strncpy(result.substring[i], input_text + i, substring_length);
    }
    return result;
}
