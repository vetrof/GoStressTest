package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("******* Передайте url *******")
	}
	url := os.Args[1]

	channel := make(chan string)

	for maxT := 0; maxT < 1000; maxT++ {
		startAllTime := time.Now()

		for i := 0; i < maxT; i += 1 {
			go requests(url, i, channel)
		}

		for i := 0; i < maxT; i += 1 {
			//fmt.Println(i, <-channel)
			_ = <-channel

		}

		fmt.Println(maxT, time.Since(startAllTime).Seconds())
	}
}

func requests(url string, i int, channel chan<- string) {
	startTime := time.Now()
	response, _ := http.Get(url)
	endTime := time.Since(startTime).Seconds()
	answer := fmt.Sprintf("%d %d %f", i, response.StatusCode, endTime)
	channel <- answer
}
