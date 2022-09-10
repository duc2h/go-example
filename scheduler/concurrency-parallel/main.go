package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)

	go g1()
	go g2()

	time.Sleep(time.Millisecond)

}

func g1() {

	time.Sleep(time.Microsecond * (1000 - 500))
	fmt.Println("g1: ")

}

func g2() {

	time.Sleep(time.Microsecond * (1000 - 20))
	fmt.Println("g2: ")

}
