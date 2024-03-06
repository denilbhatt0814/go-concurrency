package main

import(
	"fmt"
	"sync"
)

func main(){
	regularMap := make(map[int]interface{})
	// for i := 0; i < 10; i++{
	// 	regularMap[0] = i
	// } // This will throw error as in-built Maps don't support Concurrency

	// Thus we use maps from sync pkg
	syncMap := sync.Map{}

	//put
	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2
	
	// Store(key, value)
	syncMap.Store(0, 0) 
	syncMap.Store(1, 1) 
	syncMap.Store(2, 2) 

	// Get 
	regularValue, regularOk := regularMap[0]
	fmt.Println(regularValue, regularOk)
	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)
 
	// get and delete
	syncValue, loaded:= syncMap.LoadAndDelete(2)
	
	mu := sync.Mutex{}
	mu.Lock()
	regularValue = regularMap[2]
	delete(regularMap, 2)
	mu.Unlock()

	fmt.Println(syncValue, loaded, regularValue)

	// get and put
	syncValue, loaded = syncMap.LoadOrStore(1,1)

	mu.Lock()
	regularValue, regularOk = regularMap[1]
	if regularOk {
		regularMap[1] = 1
		regularValue = regularMap[1]
	}
	mu.Unlock()
	fmt.Println(syncValue, regularValue)

	// printing out stuff
	for key, value := range regularMap{
		fmt.Print(key, value, " | ")
	} 
	fmt.Println()
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Print(key, value, " | ")
		return true
	})
}