package message

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMessageParsing_PING(t *testing.T) {
	receivedString := "Command: PING\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := PING
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )
}

func TestMessageParsing_Quit(t *testing.T) {
	receivedString := "Command: QUIT\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := QUIT
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )
}

func TestMessageParsing_Who(t *testing.T) {
	receivedString := "Command: WHO\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := WHO
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )
}

func TestMessageParsing_Leave(t *testing.T) {
	receivedString := "Command: LEAVE\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := LEAVE
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )
}

func TestMessageParsing_Talk(t *testing.T) {
	receivedString := "Command: TALK\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := TALK
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )

}

func TestMessageParsing_Join(t *testing.T) {
	receivedString := "Command: JOIN\nUsername: Jeremiah\nMessage: Hello!\n\n"
	expectedCommand := JOIN
	expectedUsername := "Jeremiah"
	expectedMessage := "Hello!"

	parsedMessage := ParseMessage(receivedString)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedCommand, parsedMessage.Command)
	assert.Equal(t, expectedUsername, parsedMessage.Username )
	assert.Equal(t, expectedMessage, parsedMessage.Message )
}