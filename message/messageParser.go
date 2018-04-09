package message

import "regexp"

func ParseMessage(rcvdMessage string) *Message {
	command := getCommand(rcvdMessage)
	username := getUsername(rcvdMessage)
	message := getMessage(rcvdMessage)

	m := new(Message)
	m.Command = command
	m.Username = username
	m.Message = message

	return m
}

// Will extract the command from the received message
func getCommand(rcvdMessage string) MessageCommand {
	commandRegex := regexp.MustCompile("Command: (.*)")
	match := commandRegex.FindStringSubmatch(rcvdMessage)

	switch(match[1]) {
	case "TALK":
		return TALK
	case "JOIN":
		return JOIN
	case "LEAVE":
		return LEAVE
	default:
		return UNKNOWN
	}
}

// Retrieves the username from the received message string
func getUsername(rcvdMessage string) string {
	commandRegex := regexp.MustCompile("Username: (.*)")
	match := commandRegex.FindStringSubmatch(rcvdMessage)
	return match[1]
}

// Retrieves the actual message contents from the received string
func getMessage(rcvdMessage string) string {
	commandRegex := regexp.MustCompile("Message: (.*)")
	match := commandRegex.FindStringSubmatch(rcvdMessage)
	return match[1]
}

// TODO take the message string
// TODO pull out message
