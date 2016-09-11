package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var conf *config

// Wheel struct
type Wheel struct {
	alphabet string
	offset   int
}

// Machine struct
type Machine struct {
	wheelOrder []string
	keyphrase  string
	wheels     [3]Wheel
}

// NewMachine generator
func NewMachine(wheelOrder []string, keyphrase string) *Machine {
	keyphrase = strings.ToUpper(keyphrase)
	var wheels = [3]Wheel{}
	if len(wheelOrder) == 3 {
		i := 0
		for index := range wheelOrder {
			wheel := new(Wheel)
			wheel.alphabet, _ = getAlphabet(index, wheelOrder)
			wheels[i] = *wheel
			wheels[i].offset = strings.Index(wheels[i].alphabet, string(keyphrase[i]))
			i++
		}
	}
	return &Machine{wheelOrder, keyphrase, wheels}
}

// Open a dictionary file and return it as an array of strings.
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

// Return an very rough estimate of matched words in the message.
func findWords(words []string, message string) int {
	c := 0
	for _, search := range words {
		if strings.Count(message, search) > 0 {
			c++
		}
	}
	return c
}

// Return the a wheel alphabet,true or false.
func getAlphabet(index int, wheelOrder []string) (string, bool) {
	alphabet, ok := conf.Wheels[string(wheelOrder[index][0])][string(wheelOrder[index][1])]
	return alphabet, ok
}

func main() {

	conf = getConfig()

	wheelOrder := []string{"1A", "3A", "5A"}
	keyphrase := "XVO"

	// Validate wheel selections.
	for index, wheelName := range wheelOrder {
		if _, ok := getAlphabet(index, wheelOrder); !ok {
			log.Fatal("Error in wheel selection: ", wheelName)
		}
	}

	message := `Call me Ishmael. Some years ago—never mind how long
	precisely—having little or no money in my purse, and nothing particular to
	interest me on shore, I thought I would sail about a little and see the
	watery part of the world. It is a way I have of driving off the spleen and
	regulating the circulation. Whenever I find myself growing grim about the
	mouth; whenever it is a damp, drizzly November in my soul; whenever I find
	myself involuntarily pausing before coffin warehouses, and bringing up the
	rear of every funeral I meet; and especially whenever my hypos get such an
	upper hand of me, that it requires a strong moral principle to prevent me
	from deliberately stepping into the street, and methodically knocking
	people's hats off—then, I account it high time to get to sea as soon as I
	can. This is my substitute for pistol and ball. With a philosophical
	flourish Cato throws himself upon his sword; I quietly take to the ship.
	There is nothing surprising in this. If they but knew it, almost all men in
	their degree, some time or other, cherish very nearly the same feelings
	towards the ocean with me.`

	var words []string
	words = getDict()

	machine := *NewMachine(wheelOrder, keyphrase)
	encodedText := machine.encodeMessage(message)
	fmt.Println(encodedText)
	decodedText := machine.decodeMessage(encodedText)
	fmt.Println(decodedText)

	fmt.Println("found:", findWords(words, decodedText))

}
