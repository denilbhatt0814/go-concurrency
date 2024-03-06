package main

import(
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var missioncompleted bool
var count int = 0

func main(){
	var wg sync.WaitGroup
	wg.Add(100)

	var once sync.Once

	for i := 0; i < 100; i++{
		go func(){
			if foundTreasure(){
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()
	}
	fmt.Println("next stuff")
	wg.Wait()
	checkMissionComplition()
}

func checkMissionComplition() {
	if missioncompleted {
		fmt.Println("Mission is now completed.")
	} else {
		fmt.Println("Mission was a failure.")
	}
}

func markMissionCompleted() {
	missioncompleted = true
	fmt.Println("FOUND..!!")
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	count++
	code := rand.Intn(10)
	fmt.Println(count, code)
	return 0 == code
}