package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

var counter = 0

func main() {
	runtime.GOMAXPROCS(1)
	log.Println("running a test server")

	http.HandleFunc("/ping", ping)

	err := http.ListenAndServe(":8090", nil)
	log.Printf("server returns error: %s", err)
}

func ping(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "50" || id == "70" || id == "300" {

	} else {
		time.Sleep(time.Minute)
	}
}
