package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {

	// Get the information required to start the Sender and Receivers
	fmt.Print("Please enter your name: ")
	username := getUsernameFromUser()
	sender := new(Sender)
	receiver := new(Receiver)
	ip := net.ParseIP("255.255.255.255")

	sender.UserName = username
	sender.IPAddress = ip
	sender.Port = 11211
	receiver.Port = 11211

	// Create a channel that the Sender goroutine can use to stop the program
	stopChannel := make(chan bool)
	// Start the Sender and Receiver
	go sender.Run(stopChannel)
	go receiver.Run()
	// TODO maybe listen to these as channels, wait for them to eventually return and then quit?
	sleepUntilStopped(stopChannel)
}

// Will get the username of the user
func getUsernameFromUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	return username
}

// Will make the program try to select from nothing - this will cause the main function to do nothing while still
// allowing the other goroutines to run!
func sleepUntilStopped(stopChannel chan bool) {
	select {
	case <- stopChannel:
		return // Finally return from this function once we have received a signal to stop
	}
}