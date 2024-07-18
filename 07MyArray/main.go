package main

import (
	"fmt"
)

func main() {
	fmt.Println("These are about Arrays!")

	var AnimeList = [5]string{"Naruto", "One Piece", "Bleach", "Dragon Ball", "Vinland Saga"}

	fmt.Println("The Top 5 recomended Anime is: ", AnimeList)
	fmt.Println("Total contents in the AnimeLIst are: ", len(AnimeList))

	var WebseriesList = [4]string{"Peaky Blinders", "You", "Mr.Robot", "Suits"}

	fmt.Println("The Top 4 recomended Webseries is: ", WebseriesList)
	fmt.Println("Total contents in the WebseriesList are: ", len(WebseriesList))

}
