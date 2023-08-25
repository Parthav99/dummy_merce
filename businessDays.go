package main
//bug: negative business days
import (
	"assignmentDependencies/inputs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	holidayList := make(map[time.Time]bool)
	inputReader := bufio.NewReader(os.Stdin)

	//feeds input into the map
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := inputs.FetchInput()

		//break condition
		if holidayStr == "stop" {
			inputDate, businessDays := fetchAndValidateInput(inputReader)
			relativeDate := calculateRelativeDate(inputDate, businessDays, holidayList)
			fmt.Println("Relative Date is", relativeDate.Format("2006-01-02"))
			break
		}

		holidayDate, err := time.Parse("2006-01-02", holidayStr)
		if err != nil {
			fmt.Println("Invalid date entered. Please follow the format")
			continue
		}

		holidayList[holidayDate] = true
	}
}

//fetches and validates two input
func fetchAndValidateInput(inputReader *bufio.Reader) (time.Time, int64) {
	fmt.Print("Enter Date(YYYY-MM-DD): ")
	inputDateStr1 := inputs.FetchInput()
	fmt.Print("Enter business days: ")
	daysStr := inputs.FetchInput()

	inputDate, errDate1 := time.Parse("2006-01-02", inputDateStr1)
	businessDays, errDays := strconv.ParseInt(daysStr, 10, 64)

	if errDate1 != nil || errDays != nil {
		fmt.Println("Invalid input. Please follow the format")
		os.Exit(1)
	}

	return inputDate, businessDays
}

// Calculates relative date by adding/subtracting business days from a given date
func calculateRelativeDate(inputDate time.Time, businessDays int64, holidayList map[time.Time]bool) time.Time {
	var startDate time.Time
	var addDay int

	//increments or decrements based on business days
	if businessDays < 0 {
		addDay = -1
		businessDays = -businessDays
	} else if businessDays != 0 {
		addDay = 1
	} else {
		fmt.Println("No business days entered!")
		os.Exit(1)
	}

	for startDate = inputDate; businessDays >= 0; startDate = startDate.AddDate(0, 0, addDay) {
		if !holidayList[startDate] && startDate.Weekday() != time.Saturday && startDate.Weekday() != time.Sunday {
			businessDays--
		}
	}

	startDate = startDate.AddDate(0, 0, -addDay)

	return startDate
}

//2023-08-25
