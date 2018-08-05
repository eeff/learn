/* simplest version of calculator */

%{
#include <stdio.h>
%}

/* declare tokens */
%token NUMBER
%token ADD SUB MUL DIV AND BAR
%token LPAREN RPAREN
%token EOL

%%

calclist: /* nothing */
        | calclist exp EOL { printf("= %d\n", $2); } // EOL is end of an expression
        ;

exp: term2 /* default $$ = $1 */
   | exp AND term2 { $$ = $1 & $3; }
   | exp BAR term2 { $$ = $1 | $3; }
   ;

term2: term1 /* default $$ = $1 */
   | term2 ADD term1 { $$ = $1 + $3; }
   | term2 SUB term1 { $$ = $1 - $3; }
   ;

term1: factor /* default $$ = $1 */
    | term1 MUL factor { $$ = $1 * $3; }
    | term1 DIV factor { $$ = $1 / $3; }
    ;

factor: NUMBER /* default $$ = $1 */
      | LPAREN exp RPAREN { $$ = $2; }
      | BAR factor { $$ = $2 > 0 ? $2 : - $2; }
      ;

%%

int main(int argc, char* argv[])
{
  yyparse();
}

void yyerror(const char *s)
{
  fprintf(stderr, "error: %s\n", s);
}
