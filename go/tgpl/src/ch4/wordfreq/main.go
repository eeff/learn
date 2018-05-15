// Wordfreq reports the frequency of each word in an input text file.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: wordfreq <filename>\n")
		os.Exit(1)
	}
	freq := map[string]int{}
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		fmt.Printf("wordfreq: %s\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		freq[input.Text()]++
	}
	for w, f := range freq {
		fmt.Printf("%q : %d\n", w, f)
	}
}
