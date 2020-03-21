package main

import "fmt"

// using switch in place of if, else if, else
func main(){
	var num int
	
	switch num{
	case 2:
		fmt.Println("two")    // used in place of if
	case 4:
		fmt.Println("four")   // used in place of else if
	default:
		fmt.Println("eight")  // used in place of else
	}
}

