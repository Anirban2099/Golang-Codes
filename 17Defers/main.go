package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Defers")

	defer fmt.Println("trial") // 1st defer satement is the last one in output
	defer fmt.Println("World") // Defer used statement gets shifted after the curly brackets of func main(){ }
	fmt.Println("Hello")

	main2()

}

func main2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
