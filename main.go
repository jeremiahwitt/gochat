package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Please enter your name: ")
	username := getUsernameFromUser()
	sender := new(Sender)
	receiver := new(Receiver)
	ip := net.ParseIP("255.255.255.255")

	sender.UserName = username
	sender.IPAddress = ip
	sender.Port = 11211
	receiver.Port = 11211
	go sender.Run()
	go receiver.Run()
	// TODO maybe listen to these as channels, wait for them to eventually return and then quit?
	sleepForever()
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
func sleepForever() {
	select { }
}