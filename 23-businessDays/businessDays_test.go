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
	expected     time.Time
	holidayList  map[time.Time]bool
}{
	{"oneValue", 1, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"zeroValue", 0, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"multipleHolidays", 15, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
	}},
	{"negativeValue", -1, time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
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
