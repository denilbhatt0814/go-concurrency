package main

import(
	"fmt"
)

func main(){
	// Goroutines
	go func() {
		fmt.Println("This is the new one")
	}()
	fmt.Println("This is the old one")

}