package main

import (
	"fmt"
)

func main() {
	fmt.Println("These is about Pointers class")

	var ptr *int // Suggests that the pointer is working with integer value

	fmt.Println("The value of the pointer is", ptr)

	number1 := 77
	var ptr1 = &number1
	fmt.Println("The value of the 2nd pointer is", ptr1)   // Gives memory address
	fmt.Println("The value of the 2nd  pointer is", *ptr1) // Gives integer value

}
