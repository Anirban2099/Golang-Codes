package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is related to Switch Cases")

	rand.Seed(time.Now().UnixNano())
	//Generating random value

	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value of the dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots as well as get another chance to roll the dice")
	}

	//fallthrough statement used to run 2 cases for 1 ouput it can be written after any one case
	//it will run the next case as well
}
