package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Maps")

	CapitalCity := make(map[string]string) //Maps Syntax

	CapitalCity["MUM"] = "Mumbai"
	CapitalCity["DEL"] = "Delhi"
	CapitalCity["KOL"] = "Kolkata"

	fmt.Println("List of some Capital Cities are: ", CapitalCity)

	fmt.Println("MUM stands for: ", CapitalCity["MUM"])
	fmt.Println("DEL stands for: ", CapitalCity["DEL"])
	fmt.Println("KOL stands for: ", CapitalCity["KOL"])

	fmt.Println("Removing Delhi")
	delete(CapitalCity, "DEL")
	fmt.Println("Updated Capital City List ", CapitalCity)

	//For loop to represent values in terms of key and value

	for key, value := range CapitalCity {
		fmt.Printf("For key %v , value is %v \n", key, value)

	}
}
