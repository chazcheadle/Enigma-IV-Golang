package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "regexp"
        "strings"
  )

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

type Machine struct {
  wheel_order string
  keyphrase string
  wheels [3]Wheel
}

func NewMachine(wheel_order string, keyphrase string) *Machine {
  keyphrase = strings.ToUpper(keyphrase)
  var wheels = [3]Wheel{}
  if len(wheel_order) == 6 {
    i := 0
    for offset := 0; offset < len(wheel_order); offset+=2 {
      index := wheel_order[offset:offset+2]
      wheel := new(Wheel)
      wheel.alphabet = alphabets[index]
      wheels[i] = *wheel
      wheels[i].offset = strings.Index(wheels[i].alphabet, string(keyphrase[i]))
      i++
    }
  }
  return &Machine{wheel_order, keyphrase, wheels}
}

func (m *Machine) encode_message(message string) string {
  message = strings.ToUpper(message)
  encoder_offset := 0
  encoded_text := ""

  // Strip non-alpha characters from messge text.
  re := regexp.MustCompile("[^A-Z]")
  message = re.ReplaceAllString(message, "")

  for i := 0; i < len(message); i++ {
    // Alternate encoding wheel.
    if i % 2 == 0 {
      encoder_offset = strings.Index(m.wheels[0].alphabet, string(message[i])) - m.wheels[0].offset
    } else {
      encoder_offset = strings.Index(m.wheels[1].alphabet, string(message[i])) - m.wheels[1].offset
    }

    // Create positive offset as negative indexes won't wrap around in Go
    if encoder_offset < 0 {
      encoder_offset = len(m.wheels[2].alphabet) + encoder_offset
    }

    // Wrap the offset around the output wheel's alphabet string to simulate a physical wheel.
    if encoder_offset + m.wheels[2].offset > len(m.wheels[2].alphabet) - 1 {
      encoder_offset = encoder_offset + m.wheels[2].offset - len(m.wheels[2].alphabet)
      encoded_text += string(m.wheels[2].alphabet[encoder_offset])
    } else {
      encoded_text += string(m.wheels[2].alphabet[m.wheels[2].offset + encoder_offset])
    }

    // Add a space at every 4th letter.
    if (i+1) % 4 == 0 {
      encoded_text += " "
    }
  }
  return encoded_text
}

func (m *Machine) decode_message(message string) string {

    decoded_text := ""
    decoder_offset := 0
    j := 1

    // Strip non-alpha characters from messge text.
    re := regexp.MustCompile("[^A-Z]")
    message = re.ReplaceAllString(message, "")

    message = strings.ToUpper(message)

    for i := 0; i < len(message); i++ {
      decoder_offset = strings.Index(m.wheels[2].alphabet, string(message[i])) - m.wheels[2].offset
       if j % 2 != 0 {
        if decoder_offset + m.wheels[0].offset > len(m.wheels[0].alphabet) - 1 {
          decoder_offset = decoder_offset + m.wheels[0].offset - len(m.wheels[0].alphabet)
          decoded_text += string(m.wheels[0].alphabet[decoder_offset])
        } else if decoder_offset + m.wheels[0].offset > 0 {
          decoded_text += string(m.wheels[0].alphabet[m.wheels[0].offset + decoder_offset])
        } else {
          decoder_offset = decoder_offset + m.wheels[0].offset + len(m.wheels[0].alphabet) - 1
          decoded_text += string(m.wheels[0].alphabet[decoder_offset])
        }
      } else {
        if decoder_offset + m.wheels[1].offset > len(m.wheels[1].alphabet) - 1 {
          decoder_offset = decoder_offset + m.wheels[1].offset - len(m.wheels[1].alphabet)
          decoded_text += string(m.wheels[1].alphabet[decoder_offset])
        } else if decoder_offset + m.wheels[1].offset > 0 {
          decoded_text += string(m.wheels[1].alphabet[m.wheels[1].offset + decoder_offset])
        } else {
          decoder_offset = decoder_offset + m.wheels[1].offset + len(m.wheels[1].alphabet) - 1
          decoded_text += string(m.wheels[1].alphabet[decoder_offset])
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

func findWords(words []string, message string) int {
  c := 0
  for _, search := range words {
    if strings.Count(message, search) > 0 {
      c++;
    }
  }
  return c
}


func main() {

  wheel_order := "506070"
  keyphrase := "XVO"

  message := "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."
  encoded_text := "NVPY WKXP TZEK PPBZ CECV DPEW BDCO CCWS YGTF GYBD SLDK NSLK PETV ZSYW PSQU PKBC YFWF YKHS YZHL ICLK EDMD BUTS YWVV DUXX IYEC QFXD QKDK LUWK BDLR BCCS QRBJ SRQS GFIY MPES PVAF IUEY XUQY CVYG LKCU TKGV QKDE VVDU BMQR CABC PGXU XPEA EEXR EOCF RGDS ZSYW BMRU TKLL PKCD EDMC CWIY EUXD SUTK NSDX IYEU XFYA TKYK ZKDS RSYG WELK PMSC BAXD SWDS WVAF IUQR CZBJ QRGR CDCO CCXU XPEG EZVG DSFH PEYF ZKWT CCXD WELF IYGR CDCO CCXM XDMZ HPCY RSYO BYID QVDS PEVV IPXD STCM BCCX BMRS YAEC CRBJ LKLV YGAC XDSS YWIL QRCC CVDF RKZK DERJ YKDV PSWK CUED MKLL CXXV PYHA TKYK ZKDZ HRHL BPSK QPIX TVYJ VLCC TVYG BMWK QREU XUDK UJXC CPEP QCBD SZBC EYVC XDNS VYCU BLDK ZKYU WKRC BZMK PSAK DVQK PELU CLVS YWXD QFQR CPQC CKQV YGWK QRBG XXEY PEKD BXKS YWVK BLPK LREU LFRM QRCD XVNX BJYU XUTS SRQS WKQF SKQU BPCV EPLF BDEP XXED QRXP XPWE LJAP QSQJ QKRF DLXP QFPV YGAV PYGS QREL TSPF LFVR XXEY RYBJ DSLR NVQF QRDF GPTS WPCY RJVF YRXP LABC MSUJ XKQY HUEB CUBU TKLR XLQR CCCS LDBU TSYW LJDL DSLS YWXD QRXP XMQR CEAJ QBYK GSQV PZBP QVPY WKYS YUTK XCMK SCCK LFWK QSWK BCBU TKDX TKDS LRZK DEYK ECPE QRCP EZCM CKPS YWLU BAEC MPQR CFNK EDGS QRWK"

  var words []string
  words = getDict()

  machine := *NewMachine(wheel_order, keyphrase)
  encoded_text = machine.encode_message(message)
  fmt.Println(encoded_text)
  decoded_text := machine.decode_message(encoded_text)
  fmt.Println(decoded_text)
  fmt.Println("found:",findWords(words, decoded_text))
}



//  quitChan1 := make(chan int)
//  go decode(encoded_text, words, quitChan1)
//  quitChan2 := make(chan int)
//  go decode(encoded_text, words, quitChan2)
//  <- quitChan2
//  <- quitChan1

//  decodeCrack(encoded_text, keyphrase, wheel_order, words)
/*
func decode(encoded_text string, words []string, quitChan chan int) {
  for i := 0; i < 20; i++ {
    decoded_text := decode_message(encoded_text)
    fmt.Println(findWords(words, decoded_text))
  }
  fmt.Println("Approx words:",len(encoded_text) / 5)
  quitChan <- 1
}

func decodeCrack(encoded_text string, keyphrase string, wheel_order string, words []string) {
  // Set up machine.
  assign_wheel_order(wheel_order)
  assign_wheel_offset(keyphrase)

  decoded_text := decode_message(encoded_text)
  c := findWords(words, decoded_text)
  if len(decoded_text) / 5 < c {
    fmt.Println("key:",keyphrase,"wheels:",wheel_order)
  } else {
    fmt.Println("nope")
  }
}
*/
