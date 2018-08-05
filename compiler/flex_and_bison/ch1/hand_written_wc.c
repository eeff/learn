#include <stdio.h>

void wc(FILE *file);

int chars = 0;
int words = 0;
int lines = 0;

int main()
{
  wc(stdin);
  printf("%8d%8d%8d\n", lines, words, chars);
}

void wc(FILE *file)
{
  int is_in_a_word = 0;
  char c;
  while ((c = getchar()) != EOF)
  {
    ++chars;
    switch (c)
    {
      case 'A': case 'B': case 'C': case 'D': case 'E': case 'F': case 'G':
      case 'H': case 'I': case 'J': case 'K': case 'L': case 'M': case 'N':
      case 'O': case 'P': case 'Q': case 'R': case 'S': case 'T':
      case 'U': case 'V': case 'W': case 'X': case 'Y': case 'Z':
      case 'a': case 'b': case 'c': case 'd': case 'e': case 'f': case 'g':
      case 'h': case 'i': case 'j': case 'k': case 'l': case 'm': case 'n':
      case 'o': case 'p': case 'q': case 'r': case 's': case 't':
      case 'u': case 'v': case 'w': case 'x': case 'y': case 'z':
      {
        if ( ! is_in_a_word)
        {
          ++words;
          is_in_a_word = 1;
        }
        break;
      }
      case '\n':
      {
        ++lines;
      }
      // fall through
      default:
        is_in_a_word = 0;
    }
  }
}
