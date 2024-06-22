package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is about user input of data!!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your current location? ")

	// err syntax || comma ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Your current location is: ",input)
	fmt.Printf("Your response type is: %T \n",input)

}
