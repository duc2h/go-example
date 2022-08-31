package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	url     = "http://212.183.159.230/5MB.zip"
	workers = 5
)

type Part struct {
	Data  []byte
	Index int
}

func main() {
	var size int
	results := make(chan Part, workers)
	parts := [workers][]byte{}
	client := &http.Client{}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Headers : ", res.Header["Content-Length"])

	if header, ok := res.Header["Content-Length"]; ok {
		fileSize, err := strconv.Atoi(header[0])
		if err != nil {
			log.Fatal("File size could not be determined : ", err)
		}

		size = fileSize / workers
	}

	for i := 0; i < workers; i++ {
		go download(i, size, results)
	}

	counter := 0

	for p := range results {
		counter++

		parts[p.Index] = p.Data
		if counter == workers {
			break
		}
	}

	file := []byte{}

	// log.Println(parts)

	for i := range parts {
		file = append(file, parts[i]...)
	}

	err = ioutil.WriteFile("test.zip", file, 0700)
	if err != nil {
		log.Fatal(err)
	}
}

func download(index, size int, c chan Part) {
	client := &http.Client{}
	start := index * size
	dataRange := fmt.Sprintf("byte=%d-%d", start, start+size+1)

	// the last part
	if index == workers-1 {
		dataRange = fmt.Sprintf("byte=%d", start)
	}

	log.Println(dataRange)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// code to restart download
		return
	}

	req.Header.Add("Range", dataRange)
	res, err := client.Do(req)

	if err != nil {
		// code to restart download
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	c <- Part{Data: body, Index: index}
}
