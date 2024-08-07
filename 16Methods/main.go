package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Methods")
	User1 := User{"Anthony", "anthony2099@gmail.com", true, 20}
	//fmt.Println("Detalis of user 1: ", User1)
	//OR (more detailed output)
	fmt.Printf("Details of user 1: %+v\n", User1)

	fmt.Println("Only NAME and EMAIL needed")
	fmt.Printf("Name of user is: %v and Email of the user is: %v", User1.Name, User1.Email)
	User1.GetStatus()
}

type User struct {
	Name         string
	Email        string
	Verification bool
	Age          int
}

func (u User) GetStatus() {
	fmt.Println("Is user active? ", u.Verification)
}
