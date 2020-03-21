package main

import "fmt"

// creating own function
// simple method
func add(a int, b int) int {
	var output = a + b
	return output
}

// Compact Method
func add2(x, y int) int {
	return x + y
}
func main() {

	num := 5
	num2 := 4

	result := add(num, num2)
	result2 := add2(num, num2)
	fmt.Println(result, result2)

}
