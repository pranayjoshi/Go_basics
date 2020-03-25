package main

import "fmt"

// simple func to return an integer
func zeroval(ival int) {
	ival = 0
}

// simple func to return an *int means integer pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

// Printing results
func main() {
	i := 1
	fmt.Println("initial:", i) // initial

	zeroval(i)
	fmt.Println("zeroval:", i) // int return func

	zeroptr(&i)
	fmt.Println("zeroptr:", i) // using int pointer func

	fmt.Println("pointer:", &i) // printing the memory value using &
}
