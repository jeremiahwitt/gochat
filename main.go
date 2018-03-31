package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
	sender := new(Sender)
	go sender.Run()
	// TODO go receiver.Run()

	sleepForever()
}

// Will make the program try to select from nothing - this will cause the main function to do nothing while still
// allowing the other goroutines to run!
func sleepForever() {
	select { }
}