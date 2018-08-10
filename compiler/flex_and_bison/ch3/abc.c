/*
 * helper functions for abc.h
 */
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>
#include <math.h>
#include "abc.h"
#include "abc.tab.h"

static unsigned symhash(char *sym)
{
  unsigned int hash = 0;
  unsigned c;

  while (c = *sym++) hash = hash * 9 ^ c;

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
      sp->value = 0;
      sp->func = NULL;
      sp->syms = NULL;
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


struct ast* newast(int nodetype, struct ast *l, struct ast *r)
{
  struct ast *a = malloc(sizeof(struct ast));

  if (!a)
  {
    yyerror("out of memory");
    exit(0);
  }
  a->nodetype = nodetype;
  a->l = l;
  a->r = r;
  return a;
}

struct ast *newnum(double d)
{
  struct numval *a = malloc(sizeof(struct numval));

  if (!a)
  {
    yyerror("out of memory");
    exit(0);
  }
  a->nodetype = 'K';
  a->number = d;
  return (struct ast*) a;
}

struct ast *newcmp(int cmptype, struct ast *l, struct ast *r)
{
  struct ast *a = malloc(sizeof(struct ast));

  if (!a)
  {
    yyerror("out of memory");
    exit(1);
  }
  a->nodetype = '0' + cmptype;
  a->l = l;
  a->r = r;
  return a;
}

struct ast *newfunc(int functype, struct ast *l)
{
  struct fncall *a = malloc(sizeof(struct fncall));

  if (!a)
  {
    yyerror("out of memory");
    exit(0);
  }
  a->nodetype = 'F';
  a->l = l;
  a->functype = functype;
  return (struct ast*) a;
}

struct ast *newcall(struct symbol *s, struct ast *l)
{
  struct ufncall *a = malloc(sizeof(struct ufncall));

  if (!a)
  {
    yyerror("out of memory");
    exit(1);
  }
  a->nodetype = 'C';
  a->l = l;
  a->s = s;
  return (struct ast*) a;
}

struct ast *newref(struct symbol *s)
{
  struct symref *a = malloc(sizeof(struct symref));

  if (!a)
  {
    yyerror("out of memory");
    exit(1);
  }
  a->nodetype = 'N';
  a->s = s;
  return (struct ast*) a;
}

struct ast *newasgn(struct symbol *s, struct ast *v)
{
  struct symasgn *a = malloc(sizeof(struct symasgn));

  if (!a)
  {
    yyerror("out of memory");
    exit(1);
  }
  a->nodetype = '=';
  a->s = s;
  a->v = v;
  return (struct ast*) a;
}

struct ast *newflow(int nodetype, struct ast *cond, struct ast *tl, struct ast *el)
{
  struct flow *a = malloc(sizeof(struct flow));

  if (!a)
  {
    yyerror("out of memory");
    exit(1);
  }
  a->nodetype = nodetype;
  a->cond = cond;
  a->tl = tl;
  a->el = el;
  return (struct ast*) a;
}


void treefree(struct ast *a)
{
  switch (a->nodetype)
  {
    /* two subtrees */
  case '+':
  case '-':
  case '*':
  case '/':
  case '1': case '2': case '3': case '4': case '5': case '6':
  case 'L':
    treefree(a->r);

    /* one subtree */
  case '|':
  case 'M':
  case 'C':
  case 'F':
    treefree(a->l);

    /* no subtree */
  case 'K': case 'N':
    break;

  case '=':
    free(((struct symasgn*)a)->v);
    break;

    /* up to three subtrees */
  case 'I': case 'W':
    free(((struct flow*)a)->cond);
    if (((struct flow*)a)->tl) treefree(((struct flow*)a)->tl);
    if (((struct flow*)a)->el) treefree(((struct flow*)a)->el);
    break;

  default:
    printf("internal error: free bad node %c\n", a->nodetype);
  }
  free(a);
}

struct symlist *newsymlist(struct symbol *s, struct symlist *next)
{
  struct symlist *sl = malloc(sizeof(struct symlist));

  if (!sl)
  {
    yyerror("out of memory");
    exit(1);
  }
  sl->sym = s;
  sl->next = next;
  return sl;
}

void symlistfree(struct symlist *sl)
{
  struct symlist *nsl;
  while (sl)
  {
    nsl = sl->next;
    free(sl);
    sl = nsl;
  }
}

static double callbuiltin(struct fncall *);
static double calluser(struct ufncall *);

double eval(struct ast *a)
{
  double v;

  if (!a)
  {
    yyerror("internal error, null eval");
    exit(1);
  }

  switch (a->nodetype)
  {
    /* constant */
  case 'K':
    v = ((struct numval *)a)->number;
    break;

    /* name reference */
  case 'N':
    v = ((struct symref*)a)->s->value;
    break;

    /* assignment */
  case '=':
    v = ((struct symasgn*)a)->s->value = eval(((struct symasgn*)a)->v);
    break;

    /* expressions */
  case '+':
    v = eval(a->l) + eval(a->r);
    break;
  case '-':
    v = eval(a->l) - eval(a->r);
    break;
  case '*':
    v = eval(a->l) * eval(a->r);
    break;
  case '/':
    v = eval(a->l) / eval(a->r);
    break;
  case '|':
    v = eval(a->l);
    v = v < 0 ? -v : v;
    break;
  case 'M':
    v = -eval(a->l);
    break;

    /* comparisons */
  case '1':
    v = (eval(a->l) > eval(a->r));
    break;
  case '2':
    v = (eval(a->l) < eval(a->r));
    break;
  case '3':
    v = (eval(a->l) != eval(a->r));
    break;
  case '4':
    v = (eval(a->l) == eval(a->r));
    break;
  case '5':
    v = (eval(a->l) >= eval(a->r));
    break;
  case '6':
    v = (eval(a->l) <= eval(a->r));
    break;

    /* control flow */
    /* null expression allowed in the grammar, so check for them */

    /* if/the/else */
  case 'I':
    if (eval(((struct flow*)a)->cond))
    {
      if (((struct flow*)a)->tl)
      {
        v = eval(((struct flow*)a)->tl);
      }
      else
      {
        if (((struct flow*)a)->el)
        {
          v = eval(((struct flow*)a)->el);
        }
      }
    }
    break;

    /* while/do */
  case 'W':
    v = 0.0;
    if (((struct flow*)a)->tl)
    {
      while (eval(((struct flow*)a)->cond))
      {
        v = eval(((struct flow*)a)->tl);
      }
    }
    break;

    /* list of statements */
  case 'L':
    eval(a->l);
    v = eval(a->r);
    break;

  case 'F':
    v = callbuiltin((struct fncall*)a);
    break;

  case 'C':
    v = calluser((struct ufncall*)a);
    break;

  default:
    printf("internal error: bad node %c\n", a->nodetype);
  }

  return v;
}

static double callbuiltin(struct fncall *f)
{
  enum bifs functype = f->functype;
  double v = eval(f->l);

  switch (functype)
  {
    case B_sqrt:
      return sqrt(v);
    case B_exp:
      return exp(v);
    case B_log:
      return log(v);
    case B_print:
      printf("= %4.4g\n", v);
      return v;
    default:
      yyerror("Unknown built-in function %d", functype);
  }
  return 0.0;
}

static double calluser(struct ufncall *f)
{
  struct symbol *fn = f->s;       // function name
  struct symlist *sl = fn->syms;  // dummy arguments
  struct ast *args = f->l;        // actual arguments
  double *oldval, *newval;        // saved arg values
  double v;
  int nargs;
  int i;

  if (!fn->func)
  {
    yyerror("call to undefined function %s", fn->name);
    return 0.0;
  }

  /* count the arguments */
  for (nargs = 0; sl; sl = sl->next)
  {
    ++nargs;
  }

  /* prepare to save them */
  oldval = (double*) malloc(sizeof(double) * nargs);
  newval = (double*) malloc(sizeof(double) * nargs);
  if (!oldval || !newval)
  {
    yyerror("Out of memery in %s", fn->name);
    return 0.0;
  }

  /* evaluate the arguments */
  for (i = 0; i < nargs; ++i)
  {
    if (!args)
    {
      yyerror("too few args in call to %s", fn->name);
      free(oldval);
      free(newval);
      return 0.0;
    }

    if ('L' == args->nodetype)
    {
    // this is a list of node
      newval[i] = eval(args->l);
      args = args->r;
    }
    else
    {
    // it's the end of the list
      newval[i] = eval(args);
      args = NULL;
    }
  }

  /* save old values of dummies, assign new ones */
  sl = fn->syms;
  for (i = 0; i < nargs; ++i)
  {
    struct symbol *s = sl->sym;
    oldval[i] = s->value;
    s->value = newval[i];
    sl = sl->next;
  }

  free(newval);

  /* evaluate the function */
  v = eval(fn->func);

  /* put the real values of the dummies back */
  sl = fn->syms;
  for (i = 0; i < nargs; ++i)
  {
    struct symbol *s = sl->sym;
    s->value = oldval[i];
    sl = sl->next;
  }

  free(oldval);

  return v;
}

void dodef(struct symbol *name, struct symlist *syms, struct ast *func)
{
  if (name->syms)
  {
    symlistfree(name->syms);
  }

  if (name->func)
  {
    treefree(name->func);
  }

  name->syms = syms;
  name->func = func;
}

void yyerror(char *s, ...)
{
  va_list ap;
  va_start(ap, s);

  fprintf(stderr, "%d: error: ", yylineno);
  vfprintf(stderr, s, ap);
  fprintf(stderr, "\n");
}

int main(int argc, char* argv[])
{
  printf("> ");
  return yyparse();
}