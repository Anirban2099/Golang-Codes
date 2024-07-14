package main

import (
	"fmt"
)

func main() {
	fmt.Println("These is about Pointers class")

	var ptr *int
	// Suggests that the pointer is working with integer value

	fmt.Println("The value of the pointer is", ptr)

	Mynumber := 77
	var ptr0 = &Mynumber
	fmt.Println("The value of the 2nd pointer is", ptr0)   // Gives memory address
	fmt.Println("The value of the 2nd  pointer is", *ptr0) // Gives integer value

	*ptr0 = *ptr0 + 2
	fmt.Println("The new added +2 number is: ", ptr0)

}
