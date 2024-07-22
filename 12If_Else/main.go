package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is about IF ELSE Loops!!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your age? ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	age, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid age:", err)
		return
	}

	var result string
	if age < 18 {
		result = "Not Eligible for voting"
	} else {
		result = "Eligible for voting"
	}

	fmt.Println("So your age is:", input)
	fmt.Println("Your are:", result)
}
