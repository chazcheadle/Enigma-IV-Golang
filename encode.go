package main

import (
	"regexp"
	"strings"
)

func (m *Machine) encodeMessage(message string) string {
	message = strings.ToUpper(message)
	encoderOffset := 0
	encodedText := ""

	// Strip non-alpha characters from messge text.
	re := regexp.MustCompile("[^A-Z]")
	message = re.ReplaceAllString(message, "")

	for i := 0; i < len(message); i++ {
		// Alternate encoding wheel.
		if i%2 == 0 {
			encoderOffset = strings.Index(m.wheels[0].alphabet, string(message[i])) - m.wheels[0].offset
		} else {
			encoderOffset = strings.Index(m.wheels[1].alphabet, string(message[i])) - m.wheels[1].offset
		}

		// Create positive offset as negative indexes won't wrap around in Go
		if encoderOffset < 0 {
			encoderOffset = len(m.wheels[2].alphabet) + encoderOffset
		}

		// Wrap the offset around the output wheel's alphabet string to simulate a physical wheel.
		if encoderOffset+m.wheels[2].offset > len(m.wheels[2].alphabet)-1 {
			encoderOffset = encoderOffset + m.wheels[2].offset - len(m.wheels[2].alphabet)
			encodedText += string(m.wheels[2].alphabet[encoderOffset])
		} else {
			encodedText += string(m.wheels[2].alphabet[m.wheels[2].offset+encoderOffset])
		}

		// Add a space at every 4th letter.
		if (i+1)%4 == 0 {
			encodedText += " "
		}
	}
	return encodedText
}
