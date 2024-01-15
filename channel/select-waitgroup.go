package main

import (
	"fmt"
	"time"
)

func TestSelect1() {
	ch := make(chan int)
	ch1 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch1 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch:
			fmt.Println("v1: ", v1)
		case v2 := <-ch1:
			fmt.Println("v2: ", v2)
		}
	}

}
