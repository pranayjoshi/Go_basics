package main

import "fmt"

func main() {

	// 1. Normal Way
	// defining counter
	i := 1

	for i <= 5 {
		// Println means next line
		fmt.Println("Pranay", i) // print "Pranay" and counter no.
		i++
	}

	//  2. Linear Way
	// using semicolons to differentiate between lines
	for k := 0; k <= 5; k++ {
		fmt.Println("Pranay", k)
	}
}
