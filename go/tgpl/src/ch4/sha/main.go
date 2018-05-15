package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

// hashing algorithm
var algo = flag.String("t", "sha256", "the hash algorithm (sha256, sha384, sha512)")

func init() {
	flag.Parse()
}

func main() {
	hash := hash256
	switch *algo {
	case "sha256":
		hash = hash256
	case "sha384":
		hash = hash384
	case "sha512":
		hash = hash512
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s = %x\n", *algo, hash(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		log.Printf("sha: %v", err)
		os.Exit(1)
	}
}

func hash256(data []byte) interface{} {
	return sha256.Sum256(data)
}

func hash384(data []byte) interface{} {
	return sha512.Sum384(data)
}

func hash512(data []byte) interface{} {
	return sha512.Sum512(data)
}
