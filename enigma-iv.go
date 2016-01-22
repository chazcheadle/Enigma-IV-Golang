package main

import (
        "bufio"
        "fmt"
        "log"
        "strings"
        "os"
        "regexp"
  )
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

func encode_message(message string) string {
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
  return encoded_text
}

func decode_message(message string) string {

    decoded_text := ""
    decoder_offset := 0
    j := 1

    // Strip non-alpha characters from messge text.
    re := regexp.MustCompile("[^A-Z]")
    message = re.ReplaceAllString(message, "")

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
    return decoded_text
}

func getDict() []string {
  file, err := os.Open("words.txt")
    if err != nil {
      log.Fatal(err)
    }
  
  defer file.Close()
  
  var words []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if len(scanner.Text()) > 2 {
      words = append(words, strings.ToUpper(scanner.Text()))
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return words
}

func findWords(words []string, message string) bool {
  
  c := 0
  for _, search := range words {
    if strings.Count(message, search) > 0 { 
      c++;
    }
  }
  fmt.Println("found:",c)
  if len(message) / 5 < c {
    return true
  }
  return false
}

func main() {

  wheel_order := "506070"
  keyphrase := "XVO"
  message := "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."
  var words []string

  assign_wheel_order(wheel_order)
  assign_wheel_offset(keyphrase)

  encoded_text := encode_message(message)
//  fmt.Println(encoded_text)

  keyphrase = "XVO"
  assign_wheel_offset(keyphrase)

  decoded_text := decode_message(encoded_text)
//  fmt.Println(decoded_text)
  words = getDict()
  fmt.Println("Approx words:",len(decoded_text) / 5)

  fmt.Println(findWords(words, decoded_text))
}
