package main

import(
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	defer func(){
		fmt.Println(time.Since(start))
	}()


	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)

	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool){
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
	attacked <- true
}