package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func main() {

	url := fmt.Sprintf("http://localhost:8081/ping")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.StatusCode)
}
