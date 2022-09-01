package main

import "fmt"

func main() {
	Test()
}

func Test() {
	bufchannel := make(chan int, 1) // this is a buffered channel because it has capacity. so block main goroutine wont block if len = 1
	unbufchannel := make(chan int)  // it has not capacity, so it will block main goroutine

	fmt.Println(cap(bufchannel))   // 1
	fmt.Println(cap(unbufchannel)) // 0

	bufchannel <- 0   // its fine
	unbufchannel <- 0 // deadlock
}

// buffered channel will not block main goroutine if len <= capacity
func NoDeadlock() {
	bufchannel := make(chan int, 5) // capacity = 5

	bufchannel <- 1

	fmt.Println("lol")
}

// due to len > capacity ---> deadlock
func Deadlock() {
	bufchannel := make(chan int, 5) // capacity = 5

	bufchannel <- 1
	bufchannel <- 1
	bufchannel <- 1
	bufchannel <- 1
	bufchannel <- 1
	bufchannel <- 1

	fmt.Println("lol")
}
