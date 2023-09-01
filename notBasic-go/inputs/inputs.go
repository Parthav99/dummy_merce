package inputs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//Fetches userInput
func FetchInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	line, _ := inputReader.ReadString('\n')
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		log.Fatal("No input entered")
	}
	return line
}

//Validates dates
func FetchAndValidateDates(inputReader *bufio.Reader) (time.Time, time.Time) {
	layoutFormat := "2006-01-02"
	//loop until valid dates are entered
	for {
		fmt.Print("Enter first date(YYYY-MM-DD): ")
		inputDateStr1 := FetchInput()
		fmt.Print("Enter second date(YYYY-MM-DD): ")
		inputDateStr2 := FetchInput()

		inputDate1, err1 := time.Parse(layoutFormat, inputDateStr1)
		inputDate2, err2 := time.Parse(layoutFormat, inputDateStr2)

		if err1 != nil || err2 != nil || inputDate1.After(inputDate2) {
			fmt.Println("Invalid dates entered. Please follow the format and ensure the first date is before the second.")
			continue
		}

		return inputDate1, inputDate2
	}
}
