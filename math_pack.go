package main

import (
	"fmt"
	"math" // using math package
)

// using math.Sqrt
func main() {
	var num float64 = 4 // using
	fmt.Println("Pranay", num)
	// using Printf to print in a format
	fmt.Printf("The output is %.2f", math.Sqrt(num)) // using %.2f to print the first 2 decimal values of a float 64
}
