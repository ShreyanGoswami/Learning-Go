package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Defer is a way to delay the execution of the statement or function
	// Executed in LIFO(Last in First Out)
	fmt.Println("Start")
	defer fmt.Println("middle")
	fmt.Println("end")
	// Making a http call with defer
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // You can associate the opening of a resource to a closing of the resource right next to each otehr so that you won't forget but the closing will happen afterwards
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
