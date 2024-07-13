package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Wednesday"))

	DefinedDate := time.Date(2004, time.May, 17, 9, 17, 0, 0, time.UTC)
	fmt.Println(DefinedDate)
	fmt.Println(DefinedDate.Format("01-02-2006 Monday 15:04:05"))

}
