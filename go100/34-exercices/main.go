package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const url = "https://www.google.fr/search?q=golang"

func curl() string {
	resp, err := http.Get(url)
	if err == nil {
		return resp.Status
	}
	return "Error"
}

func main() {
	defer trackTimeElapsed(time.Now())

	unusedCh := make(chan string, 3)
	//ch <- ""
	const reqCount = 3
	for i := 0; i < reqCount; i++ {

		result := curl()
		fmt.Printf(strconv.Itoa(i) + " " + url+" responded with HTTP status %s\n", result)
	}
}

func trackTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
