/* calculator with AST */

%{
#include <stdio.h>
#include <stdlib.h>
#include "bc.h"
%}

%union {
  struct ast *a;
  double d;
}

/* declare tokens */
%token <d> NUMBER
%token EOL

%type <a> exp

%left '+' '-'
%left '*' '/'
%nonassoc '|' UMINUS

%%

calclist: /* nothing */
        | calclist exp EOL {
            printf("= %4.4g\n", eval($2));
            treefree($2);
            printf("> ");
          }
        | calclist EOL { printf("> "); } /* blank line or a comment */
        ;

exp: exp '+' exp { $$ = newast('+', $1, $3); }
   | exp '*' exp { $$ = newast('*', $1, $3); }
   | exp '-' exp { $$ = newast('-', $1, $3); }
   | exp '/' exp { $$ = newast('/', $1, $3); }
   | '|' exp     { $$ = newast('|', $2, NULL); }
   | '(' exp ')' { $$ = $2; }
   | '-' exp %prec UMINUS { $$ = newast('M', $2, NULL); }
   | NUMBER      {  $$ = newnum($1); }
   ;

%%

