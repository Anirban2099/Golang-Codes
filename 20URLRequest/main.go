package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("URL Requesting Method")
	fmt.Println(myurl)

	//Parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	Qparams := result.Query()
	fmt.Printf("The type of Query is: %T \n", Qparams)

	fmt.Println(Qparams["cws"])

	for _, val := range Qparams {
		fmt.Println("Param is: ", val)
	}

	//Creating URL method:

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "arrix.dev",
		Path:    "/golang",
		RawPath: "user=Anirban",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("Created another url: ", anotherURL)
}
