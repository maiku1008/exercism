#include "roman_numerals.h"
#include <stdlib.h>
#include <string.h>

typedef struct
{
    unsigned int arabic;
    char roman[3];
} roman_numeral;

static roman_numeral roman_numerals[] = {
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"}
};
static const unsigned int size = (sizeof(roman_numerals) / sizeof(roman_numeral));

char *to_roman_numeral(unsigned int number)
{
    char *result = calloc(10, sizeof(char));

    unsigned int i = 0;
    while (number)
    {
        if (number >= roman_numerals[i].arabic)
        {
            strcat(result, roman_numerals[i].roman);
            number -= roman_numerals[i].arabic;
        }
        else
        {
            if (i < size)
            {
                i++;
            }
        }
    }
    return result;
}
