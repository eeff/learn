CC=gcc
CFLAGS += -Wno-implicit-function-declaration

binaries=bc abc

all: bc abc

bc: bc.l bc.y bc.h bc.c
	bison -d $@.y
	flex -o $@.lex.c $@.l
	$(CC) $(CFLAGS) -o $@ $@.c $@.lex.c $@.tab.c -lfl

abc: abc.l abc.y abc.h abc.c
	bison -d $@.y
	flex -o $@.lex.c $@.l
	$(CC) $(CFLAGS) -o $@ $@.c $@.lex.c $@.tab.c -lfl -lm

clean:
	-rm -f $(binaries) *.yy.c *.tab.h *.tab.c *.out *.o *.a *.so

.PHONY:
	clean
