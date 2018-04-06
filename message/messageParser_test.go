package message

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

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