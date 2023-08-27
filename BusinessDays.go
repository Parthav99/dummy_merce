package main

//bug: negative business days
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	holidayList := make(map[time.Time]bool)
	inputReader := bufio.NewReader(os.Stdin)

	//feeds input into the map
	for {
		fmt.Print("Enter a holiday(YYYY-MM-DD) or stop: ")
		holidayStr := fetchInput()

		//break condition
		if holidayStr == "stop" {
			inputDate, businessDays := fetchAndValidateInput(inputReader)
			relativeDate := calculateEndDate(inputDate, businessDays, holidayList)
			// relativeDate := calculateRelativeDate(inputDate, businessDays, holidayList)
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

// fetches and validates two input
func fetchAndValidateInput(inputReader *bufio.Reader) (time.Time, int64) {
	fmt.Print("Enter Date(YYYY-MM-DD): ")
	inputDateStr1 := fetchInput()
	fmt.Print("Enter business days: ")
	daysStr := fetchInput()
	inputDate, errDate1 := time.Parse("2006-01-02", inputDateStr1)
	businessDays, errDays := strconv.ParseInt(daysStr, 10, 64)

	if errDate1 != nil || errDays != nil {
		fmt.Println("Invalid input. Please follow the format")
		os.Exit(1)
	}

	return inputDate, businessDays
}

func calculateEndDate(inputDate time.Time, businessDays int64, holidayList map[time.Time]bool) time.Time {

	//initializations
	counter1 := 0
	counter2 := 0
	counter3 := 0

	currentDate := inputDate
	businessDaysEntered := businessDays
	addDay := 1
	checkDay := 1

	//check if negative business days are entered
	if businessDaysEntered < 0 {
		businessDays = -businessDays
		checkDay = -1
	}

	weeks := businessDays / 5               //converts business days into whole weeks
	remainingDays := businessDays % 5       //days which remain after converting to weeks
	weekendDaysPerWeek := weeks * 2         // weekends per week
	remainingHolidays := weekendDaysPerWeek //temp storing holidays

	//iterates over remaining days and increments the weekend days
	if remainingDays > 0 {
		for i := 0; i <= int(remainingDays); i++ {
			counter1 = counter1 + 1
			if !holidayList[currentDate] && currentDate.Weekday() != time.Saturday && currentDate.Weekday() != time.Sunday {
				// continue
			} else {
				weekendDaysPerWeek++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
		}
		fmt.Println("First loop iteration", counter1)

	//if remainingDays = 0, it iterates over holidays in a week and increments the holidays
	} else {
		for remainingHolidays != 0 {
			counter2++
			if !holidayList[currentDate] {
				//continue
			} else {
				weekendDaysPerWeek++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
			remainingHolidays--
		}
	}
	fmt.Println("Second loop iteration", counter2)

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
		counter3++
		endDate = endDate.AddDate(0, 0, addDay)
	}
	fmt.Println("Third loop iteration", counter3)

	return endDate
}

// fetches input
func fetchInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	line, _ := inputReader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}

//2023-08-25

// Calculates relative date by adding/subtracting business days from a given date
// func calculateRelativeDate(inputDate time.Time, businessDays int64, holidayList map[time.Time]bool) time.Time {
// 	var startDate time.Time
// 	var addDay int
// 	iterations := 0
// 	//increments or decrements based on business days
// 	if businessDays < 0 {
// 		addDay = -1
// 		businessDays = -businessDays
// 	} else if businessDays != 0 {
// 		addDay = 1
// 	} else {
// 		fmt.Println("No business days entered!")
// 		os.Exit(1)
// 	}

// 	for startDate = inputDate; businessDays >= 0; startDate = startDate.AddDate(0, 0, addDay) {
// 		iterations++
// 		if !holidayList[startDate] && startDate.Weekday() != time.Saturday && startDate.Weekday() != time.Sunday {
// 			businessDays--
// 		}
// 	}

// 	startDate = startDate.AddDate(0, 0, -addDay)
// 	fmt.Println(iterations)
// 	return startDate
// }
