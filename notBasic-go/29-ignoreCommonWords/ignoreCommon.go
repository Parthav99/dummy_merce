// Prints the number of times each word occurs in a file. Also excludes common words

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	//accept filename input
	filename := os.Args[1]

	//open file
	file, err := os.Open(filename)
	handleError(err)
	defer file.Close()

	//read file
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	//copy words from file to a map
	wordMap := make(map[string]int)
	for fileScanner.Scan() {
		word := strings.ToLower(fileScanner.Text()) //lowercase word
		cleanedWord := wordCleaner(word)
		wordMap[cleanedWord]++
	}

	// Prints word occurrence
	commonWords := map[string]bool{
		"an": true, "the": true, "of": true, "if": true,
		"and": true, "for": true, "then": true, "to": true,
		"it": true, "or": true, "did": true, "are": true, "in": true,
		"this": true, "is": true, "what": true}

	for word, wordFreq := range wordMap {
		if len(word) < 2 || commonWords[word] {
			continue
		} else {
			fmt.Printf("Occurrence of %s:%d\n", word, wordFreq)
		}
	}
}

//Handles errors
func handleError(err error) {
	if err != nil {
		log.Fatal("\n", err)
	}
}

func wordCleaner(word string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	replacedString := reg.ReplaceAllString(word, " ")
	replacedString = strings.TrimSpace(replacedString)
	return replacedString
}
