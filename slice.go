package main

import "fmt"

func main() {

	// SLICES #1
	// creating a simple slice with indefinate value
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// SLICES #2
	// GIVING VALUES

	// Inserting values
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// assigning prior values
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// SLICES #3
	// Printing
	fmt.Println("set:", s)      // whole slice
	fmt.Println("get:", s[2])   // third element of the slice
	fmt.Println("len:", len(s)) // length of the slice

	// SLICES #4
	// appending or adding elements
	s = append(s, "d")      // sigle element
	s = append(s, "e", "f") // multiple elements
	fmt.Println("apd:", s)

	// SLICES #5
	// copying a slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	// SLICE #6
	// VIEWING SPECIFIC ELEMENTS

	// NOTE: syntax: slice[low:high]

	// from third(i.e s[2])index to the fifth(i.e s[5], where s[5] not included) index means s[2], s[3], s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// from starting means s[0] till the last element s[5], where s[5] is not included
	l = s[:5]
	fmt.Println("sl2:", l)

	// from second element i.e s[2] to the last i.e s[5]
	l = s[2:]
	fmt.Println("sl3:", l)

	// SLICES #6
	// MULTI-D SLICES

	// creating a Multi-D slices with indefinite length
	twoD := make([][]int, 3) // you should enter the maximum required value for all your slices

	// iterating in the slices
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
