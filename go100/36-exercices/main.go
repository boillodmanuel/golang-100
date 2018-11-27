package main

import (
	"fmt"
	"time"
)

const url = "https://www.google.fr/search?q=golang"

var i = 1

func curl(c chan string, e chan string) {
	if i++; i % 2 == 0{
		c <- "Success"
	} else {
		e <- "Error"
	}
}

func main() {
	defer trackTimeElapsed(time.Now())

	c := make(chan string)
	e := make(chan string)

	const reqCount = 20
	for i := 0; i < reqCount; i++ {
		go curl(c, e)
	}
	for i := 0; i < reqCount; i++ {
		select {
		case result := <-c:
			fmt.Printf(url+" OK: %s", result)
		case err := <- e:
			fmt.Printf(url+" ERROR: %s\n", err)
		}
		//result := <-c
		//fmt.Printf(url+" responded with HTTP status %s\n", result)
	}
}

func trackTimeElapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
