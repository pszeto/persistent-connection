package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func startHTTPServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(50) * time.Microsecond)
		fmt.Fprintf(w, "Hello world")
	})

	go func() {
		http.ListenAndServe(":8000", nil)
	}()

}

func startHTTPRequest() {
	url, ok := os.LookupEnv("REQUEST_URL")
	if !ok {
		log.Println("REQUEST_URL not defined.  Defaulting to http://localhost:8000")
		url = "http://localhost:8000"
	}
	iterations, ok := os.LookupEnv("INTERATIONS")
	numberOfIteration := 65535
	if !ok {
		log.Println("INTERATIONS not defined.  Defaulting to 65535")
	} else {
		i, err := strconv.Atoi(iterations)
		if err != nil {
			log.Println("INTERATIONS env varibale is not an integer.  Defaulting to 65535")
		} else {
			numberOfIteration = i
		}
	}
	counter := 0
	for i := 0; i < numberOfIteration; i++ {
		resp, err := http.Get(url)
		if err != nil {
			panic(fmt.Sprintf("Error: %v", err))
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close() // close the response body
		log.Printf("HTTP request: #%v", counter)
		log.Printf("Http request: %s - StatusCode: %d", url, resp.StatusCode)
		log.Printf("Http request: body: %s", body)
		counter += 1
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func main() {
	startHTTPServer()

	startHTTPRequest()
}
