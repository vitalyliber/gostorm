package main

import (
	"fmt"
	"net/http"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var counter = 0

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter site url: ")
	urlString, _ := reader.ReadString('\n')
	urlString = strings.TrimSuffix(urlString, "\n")

	fmt.Print("Enter number of streams: ")
	streamsString, _ := reader.ReadString('\n')
	streamsString = strings.TrimSuffix(streamsString, "\n")
	streams, _ := strconv.Atoi(streamsString)
	log.Println("Streams:", streams)

	sum := 1
	for sum < streams {
		go storm(urlString)
		log.Println("Goroutine was created with number:", sum)
		sum += 1
	}

	storm(urlString)
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
