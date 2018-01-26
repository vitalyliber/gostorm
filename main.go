package main

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	"bufio"
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
	fmt.Printf(streamsString)
	streams, _ := strconv.Atoi(streamsString)
	fmt.Println(streams)

	sum := 1
	for ; sum < streams; {
		go storm(urlString)
		fmt.Printf("Goroutine: ")
		fmt.Println(sum)
		sum += 1
	}

	storm(urlString)
}

func storm(url string) {
	for {
		fmt.Printf("Start a storm\n")
		fmt.Printf(url)
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		//defer resp.Body.Close()
		//bodyBytes, err := ioutil.ReadAll(resp.Body)
		//bodyString := string(bodyBytes)
		//fmt.Printf(bodyString)
		counter += 1
		fmt.Println(counter)
	}
}