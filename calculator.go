package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func multi(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func main() {
	// here is a small guide for the index no.
	fmt.Println("Hello Go Calculator here\n 1 = Addition\n 2 = Subtraction\n 3 = Multiplication\n 4 = Division\n")
	var inputchoice int
	fmt.Scanln(&inputchoice)
	var value1, value2 int
	fmt.Scanln(&value1, &value2)
	fmt.Println(value2)
	var result int
	switch inputchoice {
	case 1:
		fmt.Println("You Selected Addition")
		result = add(value1, value2)
		fmt.Println("Here is the result:", result)
	case 2:
		fmt.Println("You Selected Subtraction")
		result = sub(value1, value2)
		fmt.Println("Here is the result:", result)
	case 3:
		fmt.Println("You Selected Multiplication")
		result = multi(value1, value2)
		fmt.Println("Here is the result:", result)
	case 4:
		fmt.Println("You Selected Division")
		result = div(value1, value2)
		fmt.Println("Here is the result:", result)
	default:
		fmt.Println("Invalid Choice")
	}

}
