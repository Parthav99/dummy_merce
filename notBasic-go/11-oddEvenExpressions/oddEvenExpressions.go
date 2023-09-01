//Added two new cases to check even and odd numbers
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*------------- Initializations -------------*/
	var input string
	var arrayOfNumbers []float64
	var sum float64
	var operation string

	/*------------- A loop which appends input to an array -------------*/
	for {
		fmt.Print("Type proceed to exit or Enter number to be added: ")
		fmt.Scan(&input)
		numberInput, err := strconv.ParseFloat(input, 64)
		sum += numberInput

		/*------------- Break condition -------------*/
		if input == "proceed" {
			if len(arrayOfNumbers) == 0 {
				fmt.Println(operation + "Operations cannot be performed, since the input array is empty")
				os.Exit(1)
			} else {
				fmt.Print("Enter operation to be performed: ")
				fmt.Scan(&operation)
				result := choiceFunction(arrayOfNumbers, sum, operation)
				fmt.Printf("%s = %0.2f\n", operation, result)
			}
			break
		}

		/*------------- Checks if the input is a valid number before appending to an array -------------*/
		if err != nil {
			fmt.Println(input + " is not a valid number. Please enter a valid number")
		} else {
			arrayOfNumbers = append(arrayOfNumbers, numberInput)
		}
		fmt.Println(arrayOfNumbers)
	}
}

// This function performs valid operations on numbers and returns the result.
func choiceFunction(numberArray []float64, sum float64, operation string) float64 {
	var result float64
	arrLength := float64(len(numberArray))
	switch operation {
	case "count":
		result = arrLength
	case "mean":
		result = sum / arrLength
	case "min":
		result = minNumber(numberArray)
	case "max":
		result = maxNumber(numberArray)

	/*------------- Added two new operations to count even and odd numbers -------------*/
	case "counteven":
		count := 0
		for i := 0; i < len(numberArray); i++ {
			if int64(numberArray[i])%2 == 0 {
				count++
			}
		}
		result = float64(count)
	case "countodd":
		count := 0
		for i := 0; i < len(numberArray); i++ {
			if int64(numberArray[i])%2 != 0 {
				count++
			}
		}
		result = float64(count)
	default:
		fmt.Println(operation + " is an invalid Operation.")
		os.Exit(1)
	}
	return result
}

// This function returns the smallest number from an array of numbers
func minNumber(numberArray []float64) float64 {
	minNum := numberArray[0]
	for i := 0; i < len(numberArray); i++ {
		if numberArray[i] < minNum {
			minNum = numberArray[i]
		}
	}
	return minNum
}

// This function returns the largest number from an array of numbers
func maxNumber(numberArray []float64) float64 {
	maxNum := numberArray[0]
	for i := 0; i < len(numberArray); i++ {
		if numberArray[i] > maxNum {
			maxNum = numberArray[i]
		}
	}
	return maxNum
}
