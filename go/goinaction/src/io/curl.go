package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	var file io.Writer

	if len(os.Args) == 3 {
		file, err = os.Create(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		defer file.(*os.File).Close()
	} else {
		file = ioutil.Discard
	}

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
