package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Structs")
	//the alternative version of classes we use structs in Go Lang
	// no inheritance

	type User struct {
		Name         string
		Email        string
		Verification bool
		Age          int
	}

	User1 := User{"Anthony", "anthony2099@gmail.com", true, 20}
	//fmt.Println("Detalis of user 1: ", User1)
	//OR (more detailed output)
	fmt.Printf("Details of user 1: %+v\n", User1)

	fmt.Println("Only NAME and EMAIL needed")
	fmt.Printf("Name of user is: %v and Email of the user is: %v", User1.Name, User1.Email)

}
