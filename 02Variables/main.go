package main

import "fmt"

func main() {
	fmt.Printf("These are the variables codes: ")

	var x string = "Arrix"
	fmt.Println(x + " is a top gamer in India")
	fmt.Printf("this variable type is : %T \n", x)

	var y int = 10
	fmt.Println(y)
	fmt.Printf("this variable type is : %T \n", y)

	var z float64 = 99.99
	fmt.Println(z)
	fmt.Printf("this variable type is : %T \n", z)

	var w bool = true
	fmt.Println(w)
	fmt.Printf("this variable type is : %T \n", w)

	// No Var type
	a := 999999.99999 // walrus operator used i.e :=
	fmt.Println(a)
	fmt.Printf("Variable type: %T \n", a)

	//Implicit type ( where pc takes variable type upon data given )

	var r = "GOD"
	fmt.Println(r)
	fmt.Printf("Variable type is : %T \n", r)

}
