package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// to get thread max, run command `sysctl -a | grep threads-max`
const threadmax = 10

func main() {

	runtime.GOMAXPROCS(2)

	defer timeTracker(time.Now())

	wg := &sync.WaitGroup{}
	wg.Add(threadmax)
	// fmt.Println(runtime.GOMAXPROCS(0))
	// fmt.Println(runtime.NumCPU())

	for i := 0; i < threadmax; i++ {
		go g1(wg, i)

	}
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	// go g2()
	// go g3()
	// go g4()

	fmt.Println("test")

	// time.Sleep(time.Microsecond * 9999999)
}

func timeTracker(now time.Time) {
	track := time.Since(now)

	fmt.Println("finish: ", track)
}

func g1(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	// time.Sleep(time.Minute * 1)
	fmt.Println("g1: ", i)
}

func g2() {
	time.Sleep(time.Microsecond * 9994999)

	fmt.Println("g2")
}

func g3() {
	time.Sleep(time.Microsecond * 9994999)
	fmt.Println("g3")
}

func g4() {
	time.Sleep(time.Microsecond * 9994999)

	fmt.Println("g4")
}
