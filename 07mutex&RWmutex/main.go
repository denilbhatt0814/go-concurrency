package main

import(
	"fmt"
	"time"
	"sync"
)

var(
	count int
	lock sync.Mutex
	/* rwm Could be hold by multiple readers or one single writer */
	rwLock sync.RWMutex
)

func main(){
	// basics()
	readAndWrite()
}

func basics(){
	iterations := 1000
	for i := 0 ; i<iterations; i++{
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	/* Mutext allows only a single goroutine to run 
	  and access count var at a time*/
	lock.Lock()
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read(){
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println("Read locking...")
	time.Sleep(1 * time.Second)
	fmt.Println("Read Unlocking...")
}

func write(){
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking...")
	time.Sleep(1 * time.Second)
	fmt.Println("Write Unlocking...")
}