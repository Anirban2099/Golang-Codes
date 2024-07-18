package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("These are about Slices!")

	var GamesList = []string{
		"Grand Theft Auto V",
		"Sekiro",
		"God of War",
		"Elden Ring",
		"Far Cry 5",
	}

	fmt.Println("The Top recomended Games are: ", GamesList)
	fmt.Println("Total contents in the GamesLIst are: ", len(GamesList))

	fmt.Printf("The type of data in GamesList is:  %T \n ", GamesList)

	GamesList = append(GamesList, "Call of Duty", "Red Dead Redemption 2")
	//Added 2 new contents to the List making it total 7 contents
	fmt.Println("The New Top recomended Games are: ", GamesList)
	fmt.Println(" The total contents in the List now are: ", len(GamesList))

	GamesList = append(GamesList[3:4])
	/*Data before 4th content got sliced till 5th content

	other examples are: [3:] where Content shows from 4th in the List and
	[:5] Content shown till 6th content */
	fmt.Println(" Data in the 4th Content: ", GamesList)

	HighScore := make([]int, 4) //Creating Slice of 4 contents
	HighScore[0] = 100
	HighScore[1] = 300
	HighScore[2] = 500
	HighScore[3] = 700
	//HighScore[4] = 900 //Error: Index out of range
	fmt.Println("The HighScore List is: ", HighScore)

	HighScore = append(HighScore, 200, 400, 600) //Adding Values to the List

	sort.Ints(HighScore) // Soreted data in ascending order

	fmt.Println("The New and Sorted HighScore List is: ", HighScore)
	fmt.Printf("The type of data in HighScore is: %T \n", HighScore)
	fmt.Println("The total contents in HighScore are: ", len(HighScore))

	HighScore = append(HighScore[:5])
	fmt.Println("The Contents till 5th INDEX of HighScore are: ", HighScore)

}
