package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is related to Conversion Topic!! ")
	fmt.Println("Enter your DSA mark: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Your current mark of DSA is: ", input)

	// Code to convert your type of variable and adding data
	//Converting string to float

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 10 to your current mark in DSA: ", rating+10)
	}

}
