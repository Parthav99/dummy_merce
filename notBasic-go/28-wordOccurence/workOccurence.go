//Prints the number of times each word occurs in a file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//accept filename input
	filename := os.Args[1]

	//open file
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("unable to open file")
	}
	defer file.Close()

	//read file
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	//copy words from file to a map
	wordMap := make(map[string]int)
	for fileScanner.Scan() {
		word := strings.ToLower(fileScanner.Text()) //lowercase word
		wordMap[word]++
	}

	// Prints word occurence
	for word, wordFreq := range wordMap {
		fmt.Printf("Occurence of %s : %d\n", word, wordFreq)
	}

}
