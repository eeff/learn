/*
 * Declarations for a calculator
 */

/* interface to the lexer */
extern int yylineno;
void yyerror(char *s, ...);

/* nodes in abstract syntax tree */
struct ast {
  int nodetype;
  struct ast *l;
  struct ast *r;
};

struct numval {
  int nodetype;
  double number;
};

/* build an AST */
struct ast *newast(int nodetype, struct ast *l, struct ast *r);
struct ast *newnum(double d);

/* evaluate an AST */
double eval(struct ast *ast);

/* delete and free an AST */
void treefree(struct ast *ast);
