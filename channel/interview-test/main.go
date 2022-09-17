package main

import (
	"fmt"
)

func printL(ch chan int) {
	fmt.Println("so 1: ", <-ch)
	fmt.Println("so 2: ", <-ch)
}

func main() {
	ch := make(chan int)
	go printL(ch)

	ch <- 2
	ch <- 3

	// fmt.Println("len: ", len(ch))
	// fmt.Println("cap: ", cap(ch))
	fmt.Println("done")
	// time.Sleep(time.Millisecond)
}
