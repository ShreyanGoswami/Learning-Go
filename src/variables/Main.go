package main

import (
	"fmt"
	"strconv"
)

// If a variable is declared with first letter as upper case it is accessible to code outside the package
// If a variable is declared with first letter as lower case it is accessible to code inside the package

/*
Naming conventions:
The length of the variable name denotes its lifetime.
A long time means it a long lived variable
Short variables names such as i would be scoped inside a for loop/if statement etc
The most verbose variable name is for the variables at the package level
Acronyms should be UPPERCASE for eg var myHTTPRequest
*/

// To declare multiple variables at the package level
var (
	firstName string = "Shreyan"
	lastName  string = "Goswami"
	interests string = "Books, anime and games"
)

func main() {
	// Use it when you want to assign i later
	var i int
	i = 42
	// Use it for having more control on your variable types such as float32 vs float64
	var j int = 27
	// Automatic inference but gives less control
	myVariable := 42
	fmt.Println(myVariable) // Variables need to be used if they are declared
	fmt.Println(i)
	fmt.Println(j)

	// Convert integer to string
	fmt.Printf("Converted variable value %s", strconv.Itoa(myVariable))
}
