package main

import (
	"fmt"
	"math" // using math package
)

// using math.Sqrt
func main() {
	// using float64 as math.Sqrt only accept float value
	var num float64 = 4
	fmt.Println("Pranay", num)

	// using Printf to print in a format
	var result = math.Sqrt(num)

	fmt.Printf("The output is %.2f\n", result) // using %.2f to print the first 2 decimal values of a float 64

	// rouding off to the nearest
	var roundres = math.Round(result)
	fmt.Println(roundres)

	// rounding of to the top one
	var ceilres = math.Ceil(result)
	fmt.Println(ceilres)

	// rounding of to the bottom one
	var floorres = math.Floor(result)
	fmt.Println(floorres)
}
