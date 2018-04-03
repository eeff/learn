// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a real number string.
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	// handle sign
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	// decimal point
	i := strings.IndexByte(s, '.')
	if i != -1 {
		commaDigits(&buf, s[0:i])
		buf.WriteByte('.')
	}
	commaDigits(&buf, s[i+1:])
	return buf.String()
}

// comma inserts commas in a digit sequence and writes the result to a buffer.
func commaDigits(b *bytes.Buffer, s string) {
	for len(s) > 3 {
		b.WriteString(s[0:3])
		b.WriteByte(',')
		s = s[3:]
	}
	b.WriteString(s)
}
