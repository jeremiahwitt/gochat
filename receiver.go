package main

import (
	"net"
	"strconv"
	"fmt"
	"./message"
	"log"
	"strings"
)

// The Receiver receives messages from the network and displays them to the user
type Receiver struct {
	Port int
	StopChannel chan bool
	users map[string]bool
	Sender *Sender
}

func (r Receiver) Run(stopChan chan bool) {
	r.users = make(map[string]bool)
	r.StopChannel = stopChan
	listener := setupUDPListener(r)

	// Make a buffer, and then start reading messages!
	buffer := make([]byte, 2056)
	for {
		numBytes, _, err := listener.ReadFromUDP(buffer)

		// If there was an error, print it out, then listen for the next message
		if err != nil {
			fmt.Printf("Error receiving a message: %v", err)
		}

		receivedString := string(buffer[0:numBytes])
		m := message.ParseMessage(receivedString)

		switch m.Command {
		case message.JOIN:
			r.handleJoinMessage(m)
			break
		case message.TALK:
			r.handleTalkMessage(m)
			break
		case message.LEAVE:
			r.handleLeaveMessage(m)
			break
		case message.WHO:
			r.handleWhoMessage(m)
			break
		case message.QUIT:
			r.handleQuitMessage(m)
		case message.PING:
			r.handlePingMessage(m)
		}
	}
}

// Sets up the UDP Listener that will receive incoming UDP requests
func setupUDPListener(r Receiver) *net.UDPConn {
	// Resolve the address of the port
	addr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(r.Port))
	if err != nil {
		log.Fatalf("Could not listen for incoming messages with the desired port: %d", r.Port)
	}
	// Try to setup a UDPConn that we can listen to for incoming messages
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Could not listen for incoming UDP messages: %v", err)
	}
	return listener
}

// Executed by the Receiver upon receipt of a JOIN message
func (r Receiver) handleJoinMessage(m *message.Message) {
	log.Printf("%s has joined!", m.Username)

	// Add the user to the list of users who are present
	r.users[m.Username] = true
	r.Sender.SendPing()
}

// Executed by the Receiver upon receipt of a TALK message
func (r Receiver) handleTalkMessage(m *message.Message) {
	log.Printf("[%s]: %s", m.Username, m.Message)
}

// Executed by the Receiver upon receipt of a LEAVE message
func (r Receiver) handleLeaveMessage(m *message.Message) {
	log.Printf("%s has left the chat!", m.Username)

	// Remove the user from the list of all users
	delete(r.users, m.Username)
}

// Executed by the Receiver upon receipt of a WHO message. Prints
// the name of all users present
func (r Receiver) handleWhoMessage(m *message.Message) {
	allUsers := ""
	for name, _ := range r.users {
		allUsers += name + ", "
	}

	// Remove the last comma
	allUsers = strings.TrimSuffix(allUsers, ", ")
	log.Printf("Connected users: [%s]", allUsers)
}

// Executed upon receipt of QUIT message - returns via stop channel to exit program
func (r Receiver) handleQuitMessage(m *message.Message) {
	fmt.Print("Bye now!")
	r.StopChannel <- true
}

// Executed upon receipt of PING message - adds username to list of users
func (r Receiver) handlePingMessage(m *message.Message) {
	r.users[m.Username] = true
}

