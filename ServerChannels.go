package main

import "fmt"

// Main function
func main() {

	// Creating a channel using make
	mychnl := make(chan string)

	// Anonymous goroutine
	go func() {

		mychnl <- "Go"
		mychnl <- "Language"

		fmt.Println("Length of the channel is: ", len(mychnl))
		close(mychnl)
	}()

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}
}
