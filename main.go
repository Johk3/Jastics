package main

import(
	"fmt"
	"Jastics/fetch"
)


func main() {
  // Test
	// This program is mainly used to analyze, and create graphs out of
	// The user data of TextBreakerAlpha
	// Most of the work is done with Go, though sentiment analysis is done with python

	input := "/home/johk/go/src/Jastics/feedback.json"
	output := "/tmp/Jastics/"
	fmt.Println("Starting...")

	fetch.Run(input, output)
	fmt.Println("Complete!")
}
