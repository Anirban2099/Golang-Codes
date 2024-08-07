package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Functions")
	greeter()
	greeterTwo()

	result := subtracter(6, 9) // Subtracting 2 values at a time
	fmt.Println("Subtracted result is: ", result)

	result2 := adder(10, 8) // Adding 2 values at a time
	fmt.Println("Added Result is: ", result2)

	proRes := ProAdder(3, 6, 7, 9, 11, 15, 18) // Adding Many values at a time
	fmt.Println("ProAdder Result is: ", proRes)

}

func ProAdder(Values ...int) int {
	total := 0

	for _, val := range Values {
		total += val
	}

	return total
}

func adder(Val3 float64, Val4 float64) float64 {
	return Val3 + Val4
}

func subtracter(Val1 int, Val2 int) int {
	return Val2 - Val1

}
func greeter() {
	fmt.Println("Welcome User")
}

func greeterTwo() {
	fmt.Println("Welcome User 2")
}
