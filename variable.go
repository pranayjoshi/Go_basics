package main

import "fmt"

func main() {

	// declaring as integer
	var num int
	var num2 int
	// declaring value later
	num, num2 = 2, 6
	var result = num * num2
	// short ut of var by using colon
	k := 6
	// creating a constant named constant
	const cons = 9
	fmt.Print(result + k + cons)

}
