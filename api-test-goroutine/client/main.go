package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var counter = 0

func main() {
	for i := 1; i < 10000; i++ {
		go func(in int) {
			url := fmt.Sprintf("http://localhost:8090/ping?id=%d", in)
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("result %d: %d \n", in, res.StatusCode)
			counter++
		}(i)

	}

	time.Sleep(time.Second * 61)
	fmt.Println("counter: ", counter)
}
