package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int) // Creating an unbuffered channel

	go func() {
		// No sender is sending a value to the channel
		c := <-ch // Receiving a value from the channel
		fmt.Println(c)
		ch <- c + 1
	}()

	go func() {
		// No sender is sending a value to the channel
		c := <-ch // Receiving a value from the channel
		fmt.Println(c)
		ch <- c + 1
	}()

	ch <- 3

	go func() {
		c := <-ch // Receiving a value from the channel
		fmt.Println(c)
	}()

	fmt.Println("----")

	// time.Sleep(time.Second)

}

func Test() {
	bufchannel := make(chan int, 1) // this is a buffered channel because it has capacity. so main goroutine wont be blocked if bufchannel's len = 1
	unbufchannel := make(chan int)  // it has no capacity, so it will block main goroutine

	fmt.Println(cap(bufchannel))   // 1
	fmt.Println(cap(unbufchannel)) // 0

	bufchannel <- 0   // its fine
	unbufchannel <- 0 // deadlock
}

// buffered channel will not block main goroutine if len <= capacity
func NoDeadlock() {
	bufchannel := make(chan int, 5) // capacity = 5

	bufchannel <- 1

	fmt.Println("len: ", len(bufchannel))
	fmt.Println("cap: ", cap(bufchannel))

	close(bufchannel)
	value := <-bufchannel
	fmt.Println("value: ", value)
	fmt.Println("len: ", len(bufchannel))
	fmt.Println("cap: ", cap(bufchannel))

	fmt.Println("lol")
}

func PanicSendOnClosedChannel() {
	bufchannel := make(chan int, 5) // capacity = 5

	bufchannel <- 1
	bufchannel <- 2

	fmt.Println("len: ", len(bufchannel))
	fmt.Println("cap: ", cap(bufchannel))

	close(bufchannel)

	bufchannel <- 3

	fmt.Println("lol")
}

func NoPanicGetOnClosedChannel() {
	bufchannel := make(chan int, 5) // capacity = 5

	bufchannel <- 1
	bufchannel <- 2

	fmt.Println("len: ", len(bufchannel))
	fmt.Println("cap: ", cap(bufchannel))

	close(bufchannel)

	i, ok := <-bufchannel
	fmt.Println("i: ", i)
	fmt.Println("ok: ", ok)

	i1, ok1 := <-bufchannel
	fmt.Println("i1: ", i1)
	fmt.Println("ok1: ", ok1)

	<-bufchannel
	<-bufchannel
	<-bufchannel

	fmt.Println("lol")
}

func PanicSendOnClosedChannel2() {
	unbufchannel := make(chan int)

	go func() {
		unbufchannel <- 1
		unbufchannel <- 2

	}()

	<-unbufchannel
	close(unbufchannel)
	<-unbufchannel

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

func TestSelect() {
	ch := make(chan int)
	ch1 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	go func() {

		time.Sleep(time.Second)
		ch1 <- 2
	}()

	select {
	case v1 := <-ch:
		defer wg.Done()
		fmt.Println("v1: ", v1)
	case v2 := <-ch1:
		defer wg.Done()
		fmt.Println("v2: ", v2)

	}

	wg.Wait()

}
