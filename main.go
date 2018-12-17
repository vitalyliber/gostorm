package main

import (
	"flag"
	"log"
	"net/http"
)

var counter = 0
var urlString = flag.String("url", "https://some-url.com", "a site url")
var streams = flag.Int("streams", 42, "a number of streams")

func main() {
	flag.Parse()
	log.Println("Url:", *urlString)
	log.Println("Streams number:", *streams)

	sum := 1
	for sum < *streams {
		go storm(*urlString)
		log.Println("Goroutine was created with number:", sum)
		sum += 1
	}

	storm(*urlString)
}

func storm(url string) {
	for {
		log.Println("Start load test on the url:", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Got error code from url:", resp.Status)
			return
		} else {
			log.Println("Load test passed successfuly:", resp.Status)
		}

		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Println("Error while closing connection", err)
			}
			return
		}()

		counter += 1
		log.Println("Number of successful test:", counter)
	}
}
