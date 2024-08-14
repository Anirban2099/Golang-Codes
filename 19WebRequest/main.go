package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"  //WebPage 1 University of York	

func main() {
	fmt.Println("This is about Web Request Handling")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is: %T \n", response)
	fmt.Println("Response is: ", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Data type is: %T \n", data)
	fmt.Println("Data is: \n", string(data))

}
