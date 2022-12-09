package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	solution1()
	// solution2()
}

// i want the sprint only 2 line:
// line 1: ||||||||| Done line 1!
// line 2: ||||||||| Done line 2!
func problem() {
	n := 50
	wg := &sync.WaitGroup{}

	print1 := func(wg *sync.WaitGroup) {
		fmt.Print("line 1: ")
		for i := 0; i < n; i++ {
			fmt.Print("|")
			time.Sleep(time.Millisecond * 800)
		}
		fmt.Println("Done line 1!")
		wg.Done()
	}

	print2 := func(wg *sync.WaitGroup) {
		fmt.Print("line 2: ")
		for i := 0; i < n; i++ {
			fmt.Print("|")
			time.Sleep(time.Millisecond * 800)

		}
		fmt.Println("Done line 2!")
		wg.Done()
	}

	wg.Add(2)
	go print1(wg)
	go print2(wg)
	wg.Wait()
}

func solution1() {
	n := 50
	wg := &sync.WaitGroup{}
	start := time.Now()

	print1 := func(wg *sync.WaitGroup) {
		s := "line 1: "
		for i := 0; i < n; i++ {
			s += "|"
			time.Sleep(time.Millisecond * 800)
		}
		s += "Done line 1!"
		fmt.Println(s)
		wg.Done()
	}

	print2 := func(wg *sync.WaitGroup) {
		s := "line 2: "
		for i := 0; i < n; i++ {
			s += "|"
			time.Sleep(time.Millisecond * 800)

		}
		s += "Done line 2!"
		fmt.Println(s)
		wg.Done()
	}

	wg.Add(2)
	go print1(wg)
	go print2(wg)
	wg.Wait()
	end := time.Since(start).Seconds()
	fmt.Println(end)
}

func solution2() {
	n := 50
	wg := &sync.WaitGroup{}
	done := make(chan bool)
	start := time.Now()

	print1 := func(wg *sync.WaitGroup, done chan bool) {
		fmt.Print("line 1: ")
		for i := 0; i < n; i++ {
			fmt.Print("|")
			time.Sleep(time.Millisecond * 800)
		}
		fmt.Println("Done line 1!")
		wg.Done()
		done <- true
	}

	print2 := func(wg *sync.WaitGroup, done chan bool) {
		<-done
		fmt.Print("line 2: ")
		for i := 0; i < n; i++ {
			fmt.Print("|")
			time.Sleep(time.Millisecond * 800)

		}
		fmt.Println("Done line 2!")
		wg.Done()
	}

	wg.Add(2)
	go print1(wg, done)
	go print2(wg, done)
	wg.Wait()
	end := time.Since(start).Seconds()
	fmt.Println(end)
}
