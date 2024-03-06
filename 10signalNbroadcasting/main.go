package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var ready bool

func main(){
	gettingReadyForMission()
	broadcastStartOfMission()
}

func broadcastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)

	standByForMission(func(){
		fmt.Println("Ninja 1 starting mission.")
		wg.Done()
	}, beeper)

	standByForMission(func() {
		fmt.Println("Ninja 2 starting mission.")
		wg.Done()
	}, beeper)

	standByForMission(func() {
		fmt.Println("Ninja 3 starting mission.")
		wg.Done()
	}, beeper)

	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All ninjas have started their mission")
}

func standByForMission(fn func(), beeper *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()

	wg.Wait()
}

func gettingReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !ready{
		workIntervals++
	}
	cond.L.Unlock()
	fmt.Printf("We are ready! After %d work intervals.\n", workIntervals)
}

func gettingReadyForMission() {
	go gettingReady()

	workIntervals := 0
	for !ready{
		time.Sleep(5*time.Second)
		workIntervals++
	}
	fmt.Printf("We are ready! After %d work intervals.\n", workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond){
	sleep()	// Taking some time to get ready
	ready = true
	cond.Signal()
}

func gettingReady(){
	sleep()	// Taking some time to get ready
	ready = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}