package message

import (
	"html/template"
	"bytes"
	"fmt"
)

type MessageCommand string

const (
	TALK MessageCommand = "TALK"
	JOIN MessageCommand = "JOIN"
	LEAVE MessageCommand = "LEAVE"
	WHO MessageCommand = "WHO"
	QUIT MessageCommand = "QUIT"
	PING MessageCommand = "PING"
	UNKNOWN MessageCommand = "UNKNOWN"
)

// Representation of a Message as a struct
type Message struct {
	Command MessageCommand // The command to be executed by the message
	Username string // The username of the user who is talking
	Message string // The contents of the message to be ent
}

// Creates a new Message that can be used!
func BuildMessage(command MessageCommand, username string, message string) *Message {
	m := new(Message)
	m.Command = command
	m.Username = username
	m.Message = message

	return m
}

func (m *Message) String() string {

	// Add all the fields for the message to a map
	strFields := map[string]interface{}{
		"Command": m.Command,
		"UserName": m.Username,
		"Message": m.Message,
	}

	// Setup the template and the buffer for interpolation
	messageTemplate := template.Must(template.New("message").Parse(MessageFormat))
	buffer := &bytes.Buffer{}

	// Interpolate the message and make sure no errors occured
	error := messageTemplate.Execute(buffer, strFields)
	if error != nil {
		fmt.Print("Could not generate the desired message!")
		return ""
	}

	// Returns the contents of the buffer, which is the interpolated string!
	return buffer.String()
}

// Format that message contents will be interpolated into!
const MessageFormat = `Command: {{.Command}}
Username: {{.UserName}}
Message: {{.Message}}

`