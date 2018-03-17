package echo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintf(ioutil.Discard, s)
}

func echo2(args []string) {
	fmt.Fprintln(ioutil.Discard, strings.Join(args[1:], " "))
}
