package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Print("Hello world!")
	sender := new(Sender)
	receiver := new(Receiver)
	ip := net.ParseIP("255.255.255.255")

	sender.IPAddress = ip
	sender.Port = 11211
	receiver.Port = 11211
	go sender.Run()
	go receiver.Run()
	// TODO maybe listen to these as channels, wait for them to eventually return and then quit?
	sleepForever()
}

// Will make the program try to select from nothing - this will cause the main function to do nothing while still
// allowing the other goroutines to run!
func sleepForever() {
	select { }
}