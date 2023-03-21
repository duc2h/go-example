package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
			}()

			defer func() {
				close(completed)
			}()

			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	fmt.Println("doWork")
	doWork(nil)
	fmt.Println("Done.")
}

func main() {
	leak()
	leak()
	leak()

	go func() {
		http.ListenAndServe(":1234", nil)
	}()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan bool, 1)
	go func() {
		<-sigs
		done <- true
	}()
	<-done
}
