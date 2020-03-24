package main

import "fmt"

func main() {
	// Range #1 (Using it to sum up the array)
	// creating simple integer slice
	arr1 := []int{2, 3, 4, 5, 6}
	sum := 0
	for _, n := range arr1 {
		sum += n
	}
	fmt.Println(sum)

	// Range #2 ( Iterating the list to check if the found element is 5 or not)

	for _, num := range arr1 {
		if num == 5 {
			fmt.Println("5 is there in the slice")
		} else {
			fmt.Println("5 is not there in the slice")
		}
	}

	// Range #3 ( iterting in map to get keys or their values or both)

	kvs := map[string]string{"a": "apple", "b": "banana"}

	// Printing only keys
	// syntax (for looping):  for <variable to represent key> := range <map name>{<expression>}
	s := 1
	for k := range kvs {
		fmt.Println(s, "key: ", k)
		s += 1
	}

	// Printing both keys and its values
	// syntax (for looping):  for <variable to represent key>, < variable to represent value> := range <map name>{<expression>}
	for k, v := range kvs {
		fmt.Printf("%s value is %s\n", k, v)
	}

	// Range #4
	// Printing the unicode of the string
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
