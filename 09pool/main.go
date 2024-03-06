package main

import(
	"fmt"
	"sync"
)

func main() {
	// to keep track of new memory instance creation
	var numMemPieces int

	// Building the pool
	memPool := &sync.Pool{
		// in case of new allocation
		New: func() interface{} {
			numMemPieces++ // increase new instance count
			mem := make([]byte, 1) // create new memory instance
			return &mem // provide the resource
		},
	}

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++{
		go func() {
			// Checks and allots resource if available 
			mem := memPool.Get().(*[]byte) // else assigns a new one
			fmt.Println(mem)
			// Returning back the resource after use
			memPool.Put(mem) 
			wg.Done()
		}()
	}

	fmt.Printf("%d numMemPieces were created.\n", numMemPieces)
}