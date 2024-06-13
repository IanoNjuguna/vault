#include <stdio.h>

/*
 * main - calculate area of a rectangle
 * return: Success (0)
 */

int main(void)
{
  /*
   * (unsigned long int) length && width MUST be positive integers
   * (unsigned long int) 0 to 4,294,967,295
   *
   * lngth && width > 0
   */
  unsigned long int lngth;
  unsigned long int wdth;

  printf("Enter the width of the rectangle (cm): \n");
  scanf("%lu", &wdth);

  printf("Enter the length of the rectangle (cm): \n");
  scanf("%lu", &lngth);

  /*== calculate area of rectangle ==*/
  printf("The area of your rectangle is %lu cm2\n", wdth * lngth);

  return 0;
}
