/* text concordance */
%option noyywrap nodefault yylineno case-insensitive

%{
#include <ctype.h>

  struct symbol {
    char *name;
    struct ref *reflist;
  };

  struct ref {
    struct ref *next;
    char *filename;
    int flags;
    int lineno;
  };

  /* simple simtab of fixed size */
  #define NHASH 9997
  struct symbol symtab[NHASH];

  struct symbol *lookup(char*);
  void addref(int lineno, char *filename, char *word, int flags);
  void printrefs();

  char *curfilename;
%}

%%

 /* rules for concordance generator */
 /* skip common words */
a |
an |
and |
are |
as |
at |
be |
but |
for |
in |
is |
it |
of |
on |
or |
that |
the |
this |
to                /* ignore */

[a-z]+(\'(s|t))?  { addref(yylineno, curfilename, yytext, 0); }
.|\n              { /* ignore everything elese */ }

%%

int main(int argc, char* argv[])
{
  if (argc < 2) /* just read stdin */
  {
    curfilename = "(stdin)";
    yylineno = 1;
    yylex();
  }
  else
  {
    for (int i = 1; i < argc; ++i)
    {
      FILE *f = fopen(argv[i], "r");

      if (!f)
      {
        perror(argv[i]);
        return 1;
      }

      curfilename = argv[i];

      yyrestart(f);
      yylineno = 1;
      yylex();
      fclose(f);
    }
  }

  printrefs();
}

static unsigned symhash(char *sym)
{
  unsigned int hash = 0;
  unsigned c;

  while (c = tolower(*sym++)) hash = hash * 9 ^ c;

  return hash;
}

struct symbol* lookup(char* sym)
{
  struct symbol *sp = &symtab[symhash(sym)%NHASH];
  int scount = NHASH;

  while (--scount >= 0)
  {
    if (!sp->name)    /* new entry */
    {
      sp->name = strdup(sym);
      sp->reflist = 0;
      return sp;
    }
    else if (0 == strcasecmp(sp->name, sym))
    {
      return sp;
    }

    if (++sp >= symtab + NHASH)
    {
      sp = symtab;
    }
  }
  fputs("symbol table overflow\n", stderr);
  abort();
}

void addref(int lineno, char *filename, char *word, int flags)
{
  struct ref *r;
  struct symbol *sp = lookup(word);

  /* don't do dups of same line and file */
  if (sp->reflist
      && sp->reflist->lineno == lineno
      && sp->reflist->filename == filename)
  {
    return;
  }

  r = malloc(sizeof(struct ref));
  if (!r)
  {
    fputs("out of memory\n", stderr);
    abort();
  }

  r->next = sp->reflist;
  r->filename = filename;
  r->lineno = lineno;
  r->flags = flags;
  sp->reflist = r;
}

/* aux function for sorting */
static int symcompare(const void *xa, const void *xb)
{
  const struct symbol *a = xa;
  const struct symbol *b = xb;

  if (!a->name)
  {
    if (!b->name)
    {
      return 0;
    }
    return 1;
  }
  if (!b->name)
  {
    return -1;
  }
  return strcasecmp(a->name, b->name);
}

/*
 * print the references
 * sort the table alphabetically
 * the flip each entry's reflist to get it into forward order
 * and print it out
 */
void printrefs()
{
  struct symbol *sp;

  // sort the symbol table
  qsort(symtab, NHASH, sizeof(struct symbol), symcompare);

  for (sp = symtab; sp->name && sp < symtab + NHASH; ++sp)
  {
    char *prevfn = NULL;  // last printed filename, to skip dups

    // reverse the list of references
    struct ref *rp = sp->reflist;
    struct ref *rpp = 0; // previous ref
    struct ref *rpn;     // next ref

    do {
      rpn = rp->next;
      rp->next = rpp;
      rpp = rp;
      rp = rpn;
    } while (rp);

    // now print the word and its references
    printf("%10s", sp->name);
    for (rp = rpp; rp; rp = rp->next)
    {
      if (rp->filename == prevfn)
      {
        printf(" %d", rp->lineno);
      }
      else
      {
        printf(" %s:%d", rp->filename, rp->lineno);
        prevfn = rp->filename;
      }
    }
    printf("\n");
  }
}
