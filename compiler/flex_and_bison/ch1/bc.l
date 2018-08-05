/* recognize tokens for the calculator and print them out */
%{
#include "bc.tab.h"
%}

%%

"*"      { return MUL; }
"/"      { return DIV; }
"+"      { return ADD; }
"-"      { return SUB; }
"&"      { return AND;}
"|"      { return BAR; }
[0-9]+   { yylval = atoi(yytext); return NUMBER; }
"("      { return LPAREN; }
")"      { return RPAREN; }
"//".*   /* ignore comments */
\n       { return EOL; }
[ \t]    { /* ignore whitespace */ }
.        { printf("Mystery character %s\n", yytext); }

%%
