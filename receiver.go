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
	UserName string
	StopChannel chan bool
	Users map[string] *net.IP
	Sender *Sender
}

func (r Receiver) Run(stopChan chan bool) {
	r.Users = make(map[string] *net.IP)
	r.StopChannel = stopChan
	listener := setupUDPListener(r)

	// Make a buffer, and then start reading messages!
	buffer := make([]byte, 2056)
	for {
		numBytes, addr, err := listener.ReadFromUDP(buffer)

		// If there was an error, print it out, then listen for the next message
		if err != nil {
			fmt.Printf("Error receiving a message: %v", err)
		}

		receivedString := string(buffer[0:numBytes])
		m := message.ParseMessage(receivedString)

		switch m.Command {
		case message.JOIN:
			r.handleJoinMessage(m, addr)
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
			r.handlePingMessage(m, addr)
		case message.PRIVATE_TALK:
			r.handlePrivateMessage(m, addr)
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
func (r Receiver) handleJoinMessage(m *message.Message, addr *net.UDPAddr) {
	log.Printf("%s has joined!", m.Username)

	// Add the user to the list of users who are present
	r.Users[m.Username] = &addr.IP
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
	delete(r.Users, m.Username)
}

// Executed by the Receiver upon receipt of a WHO message. Prints
// the name of all users present
func (r Receiver) handleWhoMessage(m *message.Message) {
	allUsers := ""
	for name, _ := range r.Users {
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
func (r Receiver) handlePingMessage(m *message.Message, addr *net.UDPAddr) {
	r.Users[m.Username] = &addr.IP
}

// Executed when the receiver gets a private message
func (r Receiver) handlePrivateMessage(m *message.Message, addr *net.UDPAddr) {

	// If received locally, then we'll forward it out to the proper address. Otherwise, we'll print it!
	if addr.IP.Equal(net.ParseIP("127.0.0.1")) {
		if r.Users[m.Username] == nil {
			log.Printf("%s is not connected! Cannot send them a private message.", m.Username)
			return
		}

		udpAddr, err := net.ResolveUDPAddr("udp", r.Users[m.Username].String() + ":" + strconv.Itoa(r.Port))
		if err != nil {
			log.Printf("Could not send private message to %s", m.Username)
		} else {
			m.Username = r.UserName // Update the username in the message!
			SendMessageToAddress(udpAddr, m.String())
		}
	} else {
		log.Printf("[%s] (PRIVATE): %s", m.Username, m.Message)
	}
}


// Provides access to the s
func (r *Receiver) GetUserIP(username string) string {
	log.Print(r.Users)
	return r.Users[username].String()
}
