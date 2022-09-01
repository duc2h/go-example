// write a mock program download 10k record. in order to increasing speed, we will use worker pool to run .

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

const (
	records = 10000
	workers = 5
)

func main() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	mapA := make(map[int]int)
	mu := sync.Mutex{}
	queue := make(chan int, 100)
	runtime.GOMAXPROCS(5)
	procs := runtime.GOMAXPROCS(0)
	numcpu := runtime.NumCPU()

	defer timeTracker(time.Now())

	fmt.Println("cpu: ", numcpu)
	fmt.Println("procs1: ", procs)

	wg := &sync.WaitGroup{}

	wg.Add(workers + 1)

	go download(queue, wg)
	for i := 0; i < workers; i++ {
		go process(mapA, &mu, queue, i, wg)

	}

	wg.Wait()
	fmt.Println("procs3: ", procs)

	fmt.Println(mapA)
	total1 := 0
	for _, v := range mapA {
		total1 += v
	}

	fmt.Println(total1)

	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}

func timeTracker(start time.Time) {
	track := time.Since(start)

	fmt.Println("this program took ", track)
}

func appendMap(b *sync.Mutex, key int, mapA map[int]int) {
	b.Lock()
	defer b.Unlock()
	value := mapA[key]
	mapA[key] = value + 1
}

// write a function download. assign value into a queue then return it
func download(queue chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= records; i++ {
		queue <- i

		// fmt.Printf("%d has been enqueued \n", i)
	}
	close(queue)

}

// write a function receive data.
func process(m map[int]int, mu *sync.Mutex, queue <-chan int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range queue {
		appendMap(mu, index, m)
		time.Sleep(time.Microsecond * 500)

		// fmt.Printf("index: %d, value: %d \n", index, i)
	}
}
