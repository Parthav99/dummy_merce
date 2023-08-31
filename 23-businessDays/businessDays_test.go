package main

import (
	"testing"
	"time"
)

// parameterized Testing
var testcases = []struct {
	name         string
	businessDays int64
	inputTime    time.Time
	expected     time.Time //negative with holidays
	holidayList  map[time.Time]bool
}{
	{"oneValue", 1, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"sundayStart", 1, time.Date(2023, 8, 6, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 7, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"zeroValue", 0, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"multipleHolidays", 15, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 9, 21, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC): true, //check for business days as 5,10,15 --done
		time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
	}},
	{"multip", 0, time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 2, 0, 0, 0, 0, time.UTC):  true, //check for business days as 5,10,15 --done
		time.Date(2023, 8, 3, 0, 0, 0, 0, time.UTC):  true,
		time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
	}},
	{"negativeValue", -3, time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		//check for -10,-5,-7,-3
		time.Date(2023, 8, 23, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC): true,
	}},
	{"holidaysInBetween", 11, time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 17, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC):  true,
		time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 16, 0, 0, 0, 0, time.UTC): true, //check for end date as holiday --done
		time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
	}},
}

func TestCalculateEndDate(t *testing.T) {
	//iterating over slice of testcases
	for _, testCase := range testcases {
		t.Run(testCase.name, func(t *testing.T) {
			//Arrange
			expected := testCase.expected
			//Act
			got := CalculateEndDate(testCase.inputTime, testCase.businessDays, testCase.holidayList)
			//Assert
			if got != expected {
				t.Errorf("expected : %s got : %s", expected.Format("2006-01-02"), got.Format("2006-01-02"))
			}
		})
	}
}
