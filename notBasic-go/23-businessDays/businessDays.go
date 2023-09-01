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

func CalculateEndDate(inputDate time.Time, businessDays int64, holidayList map[time.Time]bool) time.Time {

	//initializations
	currentDate := inputDate
	businessDaysEntered := businessDays
	addDay := 1
	checkDay := 1
	holiday := 0

	//check if negative business days are entered
	if businessDaysEntered < 0 {
		businessDays = -businessDays
		checkDay = -1
	}

	weeks := businessDays / 5         //converts business days into business weeks
	remainingDays := businessDays % 5 //days which remain after converting to weeks
	weeksWithHolidays := (len(holidayList) + int(weeks)) / 5

	if (businessDays > 0 && weeks != 0) || (businessDays < 0 && remainingDays == 0) { //check?
		for int(weeks) <= weeksWithHolidays {
			weeks = weeks + 1 //do this conditionally
		}
	}

	weekendDaysPerWeek := weeks * 2 // weekends per week
	//iterates over remaining days and increments the weekend days
	if remainingDays > 0 {
		fmt.Println("Outside for:", currentDate)
		for i := 0; i < int(remainingDays); i++ {
			fmt.Println("Inside for:", currentDate)
			if (currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday) && currentDate != inputDate {
				fmt.Println("Inside if:", currentDate)
				weekendDaysPerWeek++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
		}
	}

	var checkEnd time.Time
	addDays := weekendDaysPerWeek + businessDays + int64(len(holidayList))
	if businessDaysEntered < 0 {
		checkEnd = inputDate.AddDate(0, 0, -int(addDays))
	} else {
		checkEnd = inputDate.AddDate(0, 0, int(addDays))
	}

	for holidays := range holidayList {
		if (inputDate.Before(holidays) && checkEnd.After(holidays)) || holidays == checkEnd {
			if holidays.Weekday() != time.Saturday && holidays.Weekday() != time.Sunday {
				holiday++
			}
		} else if holidays.Before(inputDate) && holidays.After(checkEnd) {
			if holidays.Weekday() != time.Saturday && holidays.Weekday() != time.Sunday {
				holiday++
			}
		}
	}

	fmt.Println("WeekendDays", weekendDaysPerWeek)
	fmt.Println("Holidays", holiday)

	//Calculating days to be added in order to get the end date
	totalHolidays := weekendDaysPerWeek + int64(holiday)
	totalDaysToAdd := businessDays + totalHolidays

	//if businessDays are negative, decrement days
	if businessDaysEntered < 0 {
		totalDaysToAdd = -totalDaysToAdd
		addDay = -1
	}

	//add days to the end date and iterate as long as the end date is a holiday
	endDate := inputDate.AddDate(0, 0, int(totalDaysToAdd))

	notSunday := endDate

	for holidayList[endDate] || endDate.Weekday() == time.Saturday || endDate.Weekday() == time.Sunday {
		if endDate.Weekday() == time.Sunday && notSunday.Weekday()!=time.Saturday { //condition -> Sunday, but if it is saturday then it comes to sunday and increments
			endDate = endDate.AddDate(0, 0, addDay)
		}
		endDate = endDate.AddDate(0, 0, addDay)
	}

	return endDate
}
