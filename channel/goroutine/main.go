package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go f1(ch)
	oscall := <-ch
	fmt.Println(oscall)

	time.Sleep(time.Millisecond * 200)
}

func f1(ch <-chan os.Signal) {
	<-ch

	fmt.Println("hello")
}
