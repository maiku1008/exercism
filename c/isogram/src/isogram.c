#include "isogram.h"
#include <ctype.h>
#include <string.h>

bool is_isogram(const char phrase[])
{
    if (phrase == NULL)
    {
        return false;
    }
    bool letters[26] = { false };
    for (size_t i = 0, n = strlen(phrase); i < n; i++)
    {
        char c = tolower((unsigned char) phrase[i]);
        if (!isalpha((unsigned char) c))
        {
            continue;
        }
        if (letters[c - 'a'])
        {
            return false;
        }
        letters[c - 'a'] = true;
    }
    return true;
}
