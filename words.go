package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
