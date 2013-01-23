package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	blocking := make(chan int)
	go func() { blocking <- 1 }()
	<-blocking
	fmt.Printf("Completed\n")
	// END1 OMIT
}
