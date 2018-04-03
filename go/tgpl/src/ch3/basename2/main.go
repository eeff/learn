// Basename2 reads file names from stdin and prints the base name of each one
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.og => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	// Preserve everything before last '.'
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
