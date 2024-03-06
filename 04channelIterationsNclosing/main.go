package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* METHOD-1 :- CHANNEL ITERATIONS */
// func main(){
// 	t := time.Now()
// 	defer fmt.Println(time.Since(t))
// 	channel := make(chan string)
// 	numRounds := 3
// 	go throwingNinjaStar(channel, numRounds)

// 	for i := 0; i < numRounds; i++ {
// 		fmt.Println(<-channel)
// 	}
// }

// func throwingNinjaStar(channel chan string, numRounds int){
// 	rand.Seed(time.Now().UnixNano())

// 	for i:=0; i < numRounds; i++{
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored: ", score)
// 	}
// }


/* METHOD-2 :- USING CLOSE KEYWORD */
func main(){
	t := time.Now()
	defer fmt.Println(time.Since(t))
	channel := make(chan string)
	go throwingNinjaStar(channel)

	for message := range channel{
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string){
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++{
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
	close(channel)
}