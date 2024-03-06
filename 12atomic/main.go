package main

import(
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	// this will be equivalent to 
	var mu sync.Mutex
	mu.Lock()
	sum = sum + 1
	mu.Unlock()
	fmt.Println(sum)

	// load and store
	var diffSum int64
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 1)
	fmt.Println(diffSum)

	// custom atomic concurrent ops
	var av atomic.Value
	wallace := ninja{"Wallace"}
	av.Store(wallace)
	
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		w := av.Load().(ninja)
		w.name = "Not Wallace"
		av.Store(w)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(av.Load().(ninja).name)
}

type ninja struct{
	name string
}