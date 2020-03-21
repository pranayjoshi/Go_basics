package main

import "fmt"

// Returning 2 values
// creating own addition and subraction function
// simple method
func addsub(a, b int) (output, output2 int) /*this is the returning section */ {
	output = a + b // no need to use "var" as define in the returning section
	output2 = a - b
	return
}

func main() {

	num := 6
	num2 := 3

	result, result2 := addsub(num, num2) // as add_sub returns two values we need to have to variables
	fmt.Println(result, result2)

}
