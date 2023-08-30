//Same as (21) to accepts list of holidays and two dates, then calculates work days between two dates.
package main

import (
	"assignmentDependencies/inputs"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	holidayList := make(map[time.Time]bool)
	inputReader := bufio.NewReader(os.Stdin)

	//fetching userInput(holidays) into a map(holidayList) using a loop
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := inputs.FetchInput()

		if holidayStr == "stop" {
			inputDate1, inputDate2 := inputs.FetchAndValidateDates(inputReader) //fetches and parses the dates
			weekDays := countWorkday(inputDate1, inputDate2, holidayList)       //counts number of work days
			fmt.Println("Number of working days between two dates is", weekDays)
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

func countWorkday(inputDate1 time.Time, inputDate2 time.Time, holidayList map[time.Time]bool) int {
	//initializations
	countWeekDays := 0
	hoursBetweenDates := inputDate2.Sub(inputDate1)
	totalDays := int(hoursBetweenDates.Hours() / 24) //calculating days between the two dates

	//if dates equal
	currentDate := inputDate1

	if inputDate1 == inputDate2 {
		if inputDate1.Weekday() != time.Saturday && inputDate1.Weekday() != time.Sunday &&
			!holidayList[inputDate1] {
			countWeekDays++
		}
		return countWeekDays
	}

	weekendDays := (totalDays / 7) * 2
	remainingDays := totalDays % 7

	if holidayList[currentDate] || currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	if remainingDays > 0 {
		for i := 0; i < int(remainingDays); i++ {
			if currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
				weekendDays++
			}
			currentDate = currentDate.AddDate(0, 0, 1)
		}
	}	

	totalWeekDays := totalDays - (weekendDays + len(holidayList))

	//adjust for weekends
	if totalWeekDays < 0 {
		totalWeekDays++
	}

	return totalWeekDays
}

/************************Normal Approach**************************************/
// func countWorkday(inputDate1 time.Time, inputDate2 time.Time, holidayList map[time.Time]bool) int {
// 	//initializations
// 	countWeekDays := 0
// 	hoursBetweenDates := inputDate2.Sub(inputDate1)
// 	daysBetweenDates := int(hoursBetweenDates.Hours() / 24) //calculating days between the two dates

// 	currDate := inputDate1
// 	if inputDate1 == inputDate2 {
// 		if inputDate1.Weekday() != time.Saturday && inputDate1.Weekday() != time.Sunday &&
// 			!holidayList[inputDate1] {
// 			countWeekDays++
// 		}
// 		return countWeekDays
// 	}

// 	for daysBetweenDates >= 0 {
// 		if currDate.Weekday() != time.Saturday && currDate.Weekday() != time.Sunday &&
// 			!holidayList[currDate] {
// 			countWeekDays++
// 		}

// 		daysBetweenDates--
// 		currDate = currDate.AddDate(0, 0, 1)
// 	}
// 	return countWeekDays
// }

// Calculates number of workdays between two dates
// func countWorkday(inputDate1 time.Time, inputDate2 time.Time, holidayList map[time.Time]bool) int {
// 	countWeekDays := 0
// 	for currDate := inputDate1; currDate.Before(inputDate2) || currDate.Equal(inputDate2); currDate = currDate.AddDate(0, 0, 1) {
// 		if currDate.Weekday() != time.Saturday && currDate.Weekday() != time.Sunday &&
// 			!holidayList[currDate] {
// 			countWeekDays++
// 		}
// 	}
// 	return countWeekDays
// }

/************************Logic Used**************************************/
//hoursBetweenDates := inputDate2.Sub(inputDate1) ---
// totalDays := int(hoursBetweenDates.Hours() / 24) ---
//if first date and second date
// holidays++ ----
//weekendDays := (totalDays / 7) * 2 ---
//remaining := totalDays % 7 ---
//check weekends in those remaining days
//loop over the endDate as long as it is a holiday
//holiday++
//workDays := totalDays - (holidays + weekendDays)
