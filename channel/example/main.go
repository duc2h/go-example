// write a mock program download 10k record. in order to increasing speed, we will use worker pool.

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	records = 100
	workers = 5
)

func main() {
	mapA := make(map[int]int)
	mu := sync.Mutex{}
	queue := make(chan int, 100)

	wg := &sync.WaitGroup{}

	wg.Add(workers + 1)

	go download(queue, wg)
	for i := 0; i < 5; i++ {
		go process(mapA, &mu, queue, i, wg)

	}

	wg.Wait()

	fmt.Println(mapA)
	total := 0
	for _, v := range mapA {
		total += v
	}

	fmt.Println(total)
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

		fmt.Printf("%d has been enqueued \n", i)
	}
	close(queue)

}

// write a function receive data.
func process(m map[int]int, mu *sync.Mutex, queue <-chan int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range queue {
		appendMap(mu, index, m)
		time.Sleep(time.Second)

		fmt.Printf("index: %d, value: %d \n", index, i)
	}
}
