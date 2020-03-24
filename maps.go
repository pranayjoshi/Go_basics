package main

import "fmt"

func main() {
	// syntax: make(map[key-type]val-type)

	// Creating a map
	m := make(map[string]int)

	// Assinging value
	m["k1"] = 7
	m["k2"] = 13

	// Printing whole map with keys and values
	fmt.Println("map:", m)

	// assigning m["k1"] value to a variable v1
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Printing length
	fmt.Println("len:", len(m))

	// deleting a key and its value from the map
	delete(m, "k2")
	fmt.Println("map:", m)
	// checking if it is deleted or not
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Assigning prior values
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

