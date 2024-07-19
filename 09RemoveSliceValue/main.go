package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about removing a value from slice")

	var LuckyColours = []string{
		"Black",
		"Purple",
		"Navy Blue",
		"White",
	}
	fmt.Println("All the Lucky Colours are: ", LuckyColours)
	fmt.Printf("The type of data in Lucky Colour List is:  %T \n ", LuckyColours)

	LuckyColours = append(LuckyColours, "Grey", "Violet")
	fmt.Println("the New list of Lucky Colours are: ", LuckyColours)
	fmt.Println("The total contents in the Lucky Colours List: ", len(LuckyColours))

	var index int = 3                                                      // Removing the 4th Content from the List i.e White
	LuckyColours = append(LuckyColours[:index], LuckyColours[index+1:]...) // The Proper Syntax
	fmt.Println("The New Lucky Colurs List after 4th content removed: ", LuckyColours)

}
