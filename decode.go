package main

import (
	"regexp"
	"strings"
)

func (m *Machine) decodeMessage(message string) string {

	decodedText := ""
	decoderOffset := 0
	j := 1

	// Strip non-alpha characters from messge text.
	re := regexp.MustCompile("[^A-Z]")
	message = re.ReplaceAllString(message, "")

	message = strings.ToUpper(message)

	for i := 0; i < len(message); i++ {
		decoderOffset = strings.Index(m.wheels[2].alphabet, string(message[i])) - m.wheels[2].offset
		if j%2 != 0 {
			if decoderOffset+m.wheels[0].offset > len(m.wheels[0].alphabet)-1 {
				decoderOffset = decoderOffset + m.wheels[0].offset - len(m.wheels[0].alphabet)
				decodedText += string(m.wheels[0].alphabet[decoderOffset])
			} else if decoderOffset+m.wheels[0].offset > 0 {
				decodedText += string(m.wheels[0].alphabet[m.wheels[0].offset+decoderOffset])
			} else {
				decoderOffset = decoderOffset + m.wheels[0].offset + len(m.wheels[0].alphabet) - 1
				decodedText += string(m.wheels[0].alphabet[decoderOffset])
			}
		} else {
			if decoderOffset+m.wheels[1].offset > len(m.wheels[1].alphabet)-1 {
				decoderOffset = decoderOffset + m.wheels[1].offset - len(m.wheels[1].alphabet)
				decodedText += string(m.wheels[1].alphabet[decoderOffset])
			} else if decoderOffset+m.wheels[1].offset > 0 {
				decodedText += string(m.wheels[1].alphabet[m.wheels[1].offset+decoderOffset])
			} else {
				decoderOffset = decoderOffset + m.wheels[1].offset + len(m.wheels[1].alphabet) - 1
				decodedText += string(m.wheels[1].alphabet[decoderOffset])
			}
		}
		j++
	}
	return decodedText
}
