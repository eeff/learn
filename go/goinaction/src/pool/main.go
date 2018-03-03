package main

import (
	"io"
	"log"
	"math/rand"
	"patterns/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MAX_GOROUTINES   = 25
	POOLED_RESOURCES = 2
)

var idCounter int32

type dbConnection struct {
	ID int32
}

func (c *dbConnection) Close() error {
	log.Println("Close: connection", c.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connecton", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(MAX_GOROUTINES)

	p, err := pool.New(createConnection, POOLED_RESOURCES)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < MAX_GOROUTINES; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown program")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
