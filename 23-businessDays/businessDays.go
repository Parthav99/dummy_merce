package main

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
			inputDate, businessDays := FetchAndValidateInput(inputReader)
			// relativeDate := calculateRelativeDate(inputDate, businessDays, holidayList)
			relativeDate := CalculateEndDate(inputDate, businessDays, holidayList)
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
func FetchAndValidateInput(inputReader *bufio.Reader) (time.Time, int64) {
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

//2023-08-25

func CalculateEndDate(inputDate time.Time, businessDays int64, holidayList map[time.Time]bool) time.Time {

	//initializations
	currentDate := inputDate
	businessDaysEntered := businessDays
	addDay := 1
	checkDay := 1

	//check if negative business days are entered
	if businessDaysEntered < 0 {
		businessDays = -businessDays
		checkDay = -1
	}

	weeks := businessDays / 5         //converts business days into whole weeks
	remainingDays := businessDays % 5 //days which remain after converting to weeks
	weekendDaysPerWeek := weeks * 2   // weekends per week

	//iterates over remaining days and increments the weekend days
	if remainingDays > 0 {
		for i := 0; i <= int(remainingDays); i++ {
			if holidayList[currentDate] || currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
				weekendDaysPerWeek++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
		}
	}

	//Adjusts the days to be added if start day is a Sunday
	if inputDate.Weekday() == time.Sunday {
		weekendDaysPerWeek--
	}

	//Calculating days to be added in order to get the end date
	totalHolidays := weekendDaysPerWeek
	totalDaysToAdd := businessDays + totalHolidays

	//if businessDays are negative, decrement days
	if businessDaysEntered < 0 {
		totalDaysToAdd = -totalDaysToAdd
		addDay = -1
	}

	//add days to the end date and iterate as long as the end date is a holiday
	endDate := inputDate.AddDate(0, 0, int(totalDaysToAdd))
	for holidayList[endDate] || endDate.Weekday() == time.Saturday || endDate.Weekday() == time.Sunday {
		endDate = endDate.AddDate(0, 0, addDay)
	}

	return endDate
}
