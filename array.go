package main

import "fmt"

func main() {
	// ARRAY #1
	// creating a simple integer array with the capacity of 3
	var a [3]int
	// assigning the values
	a[0] = 1
	a[1] = 2
	a[2] = 3
	// Printing array
	fmt.Println(a)

	// ARRAY #2
	// creating a string array with 4 length and assigning the values
	// syntax: <variable name> := [<length of each array>]<type>{<values>}
	ar := [4]string{"apple", "mango", "banana", "cherry"}
	// Printing array in a format
	fmt.Printf("I love %s, I hate %s, I like %s, I rarely eat %sn", ar[1], ar[0], ar[2], ar[3])

	// ARRAY #3
	// creting an integer array with indefinite length and assigning the values
	arr := [...]int{1, 5, 6, 0}
	fmt.Println(arr)
	// example by changing values
	arr[1] = 2
	fmt.Println(arr)
	// Printing specific values
	fmt.Println(arr[2])
	// Printing length
	fmt.Println(len(arr))
	// ARRAY #4
	// creating a 2D array syntax: <variable name> := [<amount of array>][<length of each array>]<type>{<values>}
	arra := [2][3]int{{1, 2, 3}, {3, 2, 1}} // assinging values
	// Printing array and length
	fmt.Println(arra, len(arra))

	// ARRAY #5
	// creating a 2D array and entering values seprately
	var array [2][3]int
	array[0][0] = 100
	array[0][1] = 200
	array[0][2] = 300
	array[1][0] = 400
	array[1][1] = 500
	array[1][2] = 600

	// Accessing the values of the array by Iterating the 2D array
	fmt.Println("Elements of array 2")
	for p := 0; p < 3; p++ {
		for q := 0; q < 3; q++ {
			fmt.Println(array[p][q])

		}
	}
}
