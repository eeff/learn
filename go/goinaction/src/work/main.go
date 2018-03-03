package main

import (
	"log"
	"patterns/work"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Work() {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	if p == nil {
		log.Println("Worker pool creation failure")
		return
	}

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
