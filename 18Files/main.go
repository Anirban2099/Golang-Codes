package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is about Files Handling")

	content := "Learn about Berserk anime from tis link:  - https://digital.darkhorse.com/pages/156/berserk"
	file, err := os.Create("./BerserkInfo.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content) //gives the info about length of the file

	if err != nil {
		panic(err)
	}

	fmt.Println("lenth is: ", length)
	defer file.Close()

	ReadtheFile("./BerserkInfo.txt")

}

func ReadtheFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content inside the file: \n", data)
	fmt.Println("Content inside the file in string type: \n", string(data))

}
