package main

import "fmt"
import "strings"
import "regexp"

var wheels = []*Wheel{}

type Wheel struct {
  alphabet string
  offset int
}

var alphabets = map[string]string {
  "50": "XGBRCJSQIEFTVHYAPOWZNULKMD",
  "51": "REHKVMQTFSJNXBWZGDOALCUPIY",
  "52": "SUBWDVRFMKHPOLZCGXINQAJEYT",
  "53": "YOBEZALKIHRCUFVQWTSMPXGNJD",
  "60": "VIWNXUPTCRHJMBZYAKDOLQSEGF",
  "61": "DUSYOCQGZALBKFWHJIVEMPXRNT",
  "62": "DASQOPELGKUVBTWYRCINHMXJFZ",
  "63": "ZFTIKGOPJLYUDHNMAWVSRECXBQ",
  "70": "OSADNJLUXCRQZTHEVBGFYIPKWM",
  "71": "INFEGJBTMPZSQWUYKRXHCDLVOA",
  "72": "OZBNXIALJFRWGKQCDVYMTEUSHP",
  "73": "XGWMOVIZDEFYSPBRTJHAQCKULN",
}

func assign_wheel_order(wheel_order string) {
  if len(wheel_order) == 6 {
    for offset := 0; offset < len(wheel_order); offset+=2 {
      index := wheel_order[offset:offset+2]
      wheel := new(Wheel)
      wheel.alphabet = alphabets[index]
      wheels = append(wheels, wheel)
    }
  }
}

func assign_wheel_offset(keyphrase string) {
  keyphrase = strings.ToUpper(keyphrase)
  for i := 0; i < len(wheels); i++ {
    wheels[i].offset = strings.Index(wheels[i].alphabet, string(keyphrase[i]))
  }
}

func encode_message(message string) {
  message = strings.ToUpper(message)
  encoder_offset := 0
  encoded_text := ""

  // Strip non-alpha characters from messge text.
  re := regexp.MustCompile("[^A-Z]")
  message = re.ReplaceAllString(message, "")

  for i := 0; i < len(message); i++ {
    // Alternate encoding wheel.
    if i % 2 == 0 {
      encoder_offset = strings.Index(wheels[0].alphabet, string(message[i])) - wheels[0].offset
    } else {
      encoder_offset = strings.Index(wheels[1].alphabet, string(message[i])) - wheels[1].offset
    }

    // Create positive offset as negative indexes won't wrap around in Go
    if encoder_offset < 0 {
      encoder_offset = len(wheels[2].alphabet) + encoder_offset
    }

    // Wrap the offset around the output wheel's alphabet string to simulate a physical wheel.
    if encoder_offset + wheels[2].offset > len(wheels[2].alphabet) - 1 {
      encoder_offset = encoder_offset + wheels[2].offset - len(wheels[2].alphabet)
      encoded_text += string(wheels[2].alphabet[encoder_offset])
    } else {
      encoded_text += string(wheels[2].alphabet[wheels[2].offset + encoder_offset])
    }

    // Add a space at every 4th letter.
    if (i+1) % 4 == 0 {
      encoded_text += " "
    }
  }
  fmt.Println(message)
  fmt.Println(encoded_text)
}

func decode_message(message string) {

    decoded_text := ""
    decoder_offset := 0
    j := 1
    message = "QUXFXFENCFQCCFLTSY"
    message = strings.ToUpper(message)

      for i := 0; i < len(message); i++ {
        decoder_offset = strings.Index(wheels[2].alphabet, string(message[i])) - wheels[2].offset

        if j % 2 != 0 {
          if decoder_offset + wheels[0].offset > len(wheels[0].alphabet) - 1 {
            decoder_offset = decoder_offset + wheels[0].offset - len(wheels[0].alphabet)
            decoded_text += string(wheels[0].alphabet[decoder_offset])
          } else if decoder_offset + wheels[0].offset > 0 {
            decoded_text += string(wheels[0].alphabet[wheels[0].offset + decoder_offset])
          } else {
            decoder_offset = decoder_offset + wheels[0].offset + len(wheels[0].alphabet) - 1
            decoded_text += string(wheels[0].alphabet[decoder_offset])
          }
        } else {
          if decoder_offset + wheels[1].offset > len(wheels[1].alphabet) - 1 {
            decoder_offset = decoder_offset + wheels[1].offset - len(wheels[1].alphabet)
            decoded_text += string(wheels[1].alphabet[decoder_offset])
          } else if decoder_offset + wheels[1].offset > 0 {
            decoded_text += string(wheels[1].alphabet[wheels[1].offset + decoder_offset])
          } else {
            decoder_offset = decoder_offset + wheels[1].offset + len(wheels[1].alphabet) - 1
            decoded_text += string(wheels[1].alphabet[decoder_offset])
          }
        }

        j += 1
      }
    fmt.Println(message)
    fmt.Println(decoded_text)
}

func main() {

  wheel_order := "506070"
  keyphrase := "EMC"
  message := "This is a test message"

  assign_wheel_order(wheel_order)
  assign_wheel_offset(keyphrase)
  encode_message(message)
  decode_message(message)
}
