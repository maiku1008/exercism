#include "resistor_color.h"

int color_code(resistor_band_t color)
{
    return color;
}

int * colors(void)
{
    static int r[10];
    for (int i = BLACK; i <= WHITE; i++)
    {
        r[i] = i;
    }
    return r;
}
