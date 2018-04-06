package main

import (
	"net"
	"strconv"
	"fmt"
	"./message"
)

// The Receiver receives messages from the network and displays them to the user
type Receiver struct {
	Port int
}

func (r Receiver) Run() {
	// Resolve the address of the port
	addr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(r.Port))

	if err != nil {
		fmt.Errorf("Could not listen for incoming messages with the desired port: %d", r.Port)
	}

	// Try to setup a UDPConn that we can listen to for incoming messages
	listener, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Errorf("Could not listen for incoming UDP messages: %v", err)
	}

	// Make a buffer, and then start reading messages!
	buffer := make([]byte, 2056)
	for {
		numBytes, senderAddr, err := listener.ReadFromUDP(buffer)

		// If there was an error, print it out, then listen for the next message
		if err != nil {
			fmt.Printf("Error receiving a message: %v", err)
		}

		receivedString := string(buffer[0:numBytes])
		message := message.ParseMessage(receivedString)
		fmt.Println("Got string: " + message.Message + " from " + message.Username) // TODO remove
		fmt.Println(senderAddr) // TODO remove
		// TODO generateMessage(receivedString, senderAddr)
		// TODO printMessage()

	}

}

