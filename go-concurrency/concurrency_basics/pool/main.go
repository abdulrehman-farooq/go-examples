package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	myPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := myPool.Get().(*[]byte)
			defer myPool.Put(mem)

			// Assume something interesting, but quick is being done with
			// this memory.
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
