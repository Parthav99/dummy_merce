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
	counter := 0
	currentDate := inputDate
	businessDaysEntered := businessDays
	addDay := 1
	checkDay := 1
	if businessDays < 0 {
		businessDays=-businessDays
		checkDay=-1
	}
	
	fullWeeks := businessDays / 5
	remainingDays := businessDays % 5
	weekendsInFullWeeks := fullWeeks * 2
	remainHolidays := weekendsInFullWeeks

	if remainingDays > 0 {
		for i := 0; i <= int(remainingDays); i++ {
			counter = counter + 1
			if !holidayList[currentDate] && currentDate.Weekday() != time.Saturday && currentDate.Weekday() != time.Sunday {
				// continue
			} else {
				weekendsInFullWeeks++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
		}
		fmt.Println("First loop iteration", counter)
	} else {
		for remainHolidays != 0 {
			counter++
			if !holidayList[currentDate] {
				//continue
			} else {
				fmt.Println(currentDate)
				weekendsInFullWeeks++
			}
			currentDate = currentDate.AddDate(0, 0, checkDay)
			remainHolidays--
		}
	}
	fmt.Println("Second loop iteration", counter)

	// fmt.Println("Current Date:", currentDate)
	if inputDate.Weekday() == time.Sunday {
		weekendsInFullWeeks--
	}

	totalWeekends := weekendsInFullWeeks
	totalDaysToAdd := businessDays + totalWeekends
	if businessDaysEntered < 0 {
		totalDaysToAdd = -totalDaysToAdd
	}
	endDate := inputDate.AddDate(0, 0, int(totalDaysToAdd))

	if businessDaysEntered < 0 {
		addDay = -1
	}

	counter2 := 0
	for holidayList[endDate] || endDate.Weekday() == time.Saturday || endDate.Weekday() == time.Sunday {
		counter2++
		endDate = endDate.AddDate(0, 0, addDay)
	}
	fmt.Println("Third loop iteration", counter2)
	// endDate = endDate.AddDate(0, 0, addDay)
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
