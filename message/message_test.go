package message

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Tests that the string is generated as expected
func TestMessageString(t *testing.T) {
	command := TALK
	username := "Jeremiah"
	message := "Hello!"

	generatedMessage := BuildMessage(command, username, message)
	generatedString := generatedMessage.String()

	expectedString := "Command: TALK\nUsername: Jeremiah\nMessage: Hello!\n\n"

	assert.Equal(t, expectedString, generatedString)
}

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}