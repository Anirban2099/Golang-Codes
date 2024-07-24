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

	for i := 0; i < len(Pokemon); i++ {
		fmt.Println(Pokemon[i])
	}

	
}
