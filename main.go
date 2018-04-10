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

	// Create channels that the Sender and Receivergoroutine can use to stop the program
	senderStopChannel := make(chan bool)
	receiverStopChannel := make(chan bool)
	// Start the Sender and Receiver
	go sender.Run(senderStopChannel)

	go receiver.Run(receiverStopChannel)
	// TODO maybe listen to these as channels, wait for them to eventually return and then quit?
	sleepUntilStopped(senderStopChannel, receiverStopChannel)
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
func sleepUntilStopped(senderStopChannel chan bool, receiverStopChannel chan bool) {
	select {
	case <- senderStopChannel:
		// DO nothing - we now will wait for the receiver
	}

	select {
	case <- receiverStopChannel:
		return
	}
}