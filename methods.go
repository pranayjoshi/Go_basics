package main

import "fmt"

// methods defined on the struct types
// creating a struct named rect which has width and height, 2 integer field
type rect struct {
	width, height int
}

// creting a method to tell the are and perimeter

// method r.area
func (r *rect) area() int {
	return r.width * r.height
}

//method r.perim
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	// prints out the area and perimeter of the above values in the fields
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	// using pointers
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
