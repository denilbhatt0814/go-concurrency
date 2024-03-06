package main

import(
	"fmt"
)

func main(){
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	/* This wouldn't let us elect our CAP*/
	// fmt.Println(<-ninja1)
	// fmt.Println(<-ninja2)

	/*
		Thus we use select keyword:
		it waits till goroutine to complete 
		and runs case on random
	*/

	select{
		case message := <-ninja1:
			fmt.Println(message)
		case message := <-ninja2:
			fmt.Println(message)
	}
	roughlyFair() // Shows it has probability of 50%
}

func captainElect(ninja chan string, message string){
	ninja<-message
}

func roughlyFair(){
	ninja1 := make(chan interface{}); close(ninja1)
	ninja2 := make(chan interface{}); close(ninja2)

	var ninja1count, ninja2count int
	for i:=0; i<1000; i++{
		select{
		case <- ninja1:
			ninja1count++
		case <- ninja2:
			ninja2count++
		}
	}

	fmt.Printf("ninja1count: %d, ninja2count: %d\n", ninja1count, ninja2count)
}