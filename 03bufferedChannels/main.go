package main

import "fmt"

func main(){
	// CASE-1:
	// channel := make(chan string)
	// go func(){
	// 	channel <- "first message"
	// }()
	// fmt.Println(<-channel)
	/* This works as sending the msg in channels occurs 
		in seperate goroutine which runs along side main
		Thus when msg is sent, at same msg could be recieved.
	*/

	// CASE-2:
	// channel := make(chan string)
	// channel <- "first meassage"
	// fmt.Println(<-channel)
	/* This case goes in deadlock condn.
		as when the msg is sent in Channel it isn't received at 
		SAME TIME since both runs on same main goroutine
	*/

	// CASE-3: Using buffer channel 
	channel := make(chan string, 1) 
	// Providing a capacity of 1 to Chan makes it buffer chan for 1 input
	channel <- "first message"
	fmt.Println(<- channel)

	// Capacity could also be increased 
	ch := make(chan string, 2)
	ch <- "message 1"
	ch <- "message 2"
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	/*IMP :- 
		Channels are First in first out sys (message queue)
		thus what was entered 1st will come out 1st
	*/
}