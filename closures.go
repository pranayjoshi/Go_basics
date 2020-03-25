package main

import "fmt"

// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

// This function test returns another function, which we define anonymously in the body of test. The returned function closes over the variable i to form a closure.
func test() func() int {
	i := 101
	return func() int {
		i--
		return i
	}
}

func main() {
	// Printing back counting from 100 to 90
	test1 := test()
	for i := 1; i <= 10; i++ {
		fmt.Println(test1())
	}
	// it refreshes every time
	test2 := test()
	fmt.Println(test2())
}
