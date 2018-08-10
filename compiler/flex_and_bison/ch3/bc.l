/* recognize tokens for the calculator */
%option noyywrap nodefault
%{
#include "bc.h"
#include "bc.tab.h"
%}

/* float exponent */
EXP   ([Ee][-+]?[0-9]+)

%%

"*" |
"/" |
"+" |
"-" |
"&" |
"|" |
"(" |
")" { return yytext[0]; }

[0-9]+"."[0-9]*{EXP}? |
"."?[0-9]+{EXP}?      { yylval.d = atof(yytext); return NUMBER; }

\n       { return EOL; }
"//".*   /* ignore comments */
[ \t]    { /* ignore whitespace */ }
.        { yyerror("Mystery character %s\n", *yytext); }


%%
