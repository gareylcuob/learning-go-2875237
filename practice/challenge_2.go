// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	value1Float, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	// Convert the second string to a float64
	value2Float, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	// Calculate and return the result

	resultFloat := value1Float + value2Float

	fmt.Println(resultFloat)

	return 0
}
