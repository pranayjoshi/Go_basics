package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// NewPerson constructs a new person struct with the given name
// return a pointer to local variable as a local variable will survive the scope of the function.
func NewPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
