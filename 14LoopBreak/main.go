package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Loops Break")

	Pokemon := []string{ //Slice used
		"Charizard",
		"Venusaur",
		"Blastoise",
		"Raichu",
	}

	fmt.Println("list of Pokemons: ", Pokemon)

	//Process:1
	for i := 0; i < len(Pokemon); i++ {
		fmt.Println(Pokemon[i])
	}
	//Process:2
	for p := range Pokemon {
		fmt.Println(Pokemon[p])
	}
	//Process:3 (Index added to contents) if index is not needed then replace index with "_"
	for index, Pokemon := range Pokemon {
		fmt.Printf("Index is %v and value is %v\n", index, Pokemon)
	}

	rougeValue := 1

	for rougeValue < 10 {
		fmt.Println("Value is : ", rougeValue)
		rougeValue++
	}
}
