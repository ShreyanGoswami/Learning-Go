package main

import "fmt"

func main() {
	// Simplest - curly braces are mandatory
	if true {
		fmt.Println("Success!")
	}
	// Initialiser if statement
	myMap := map[string]int{
		"val1": 1,
		"val2": 2,
	}
	if _, ok := myMap["val1"]; ok {
		fmt.Println("Found key in map")
	}
	// Switch statements - no breaks only one condition and multiple values in each case
	switch 5 {
	case 1, 5, 10: // can have also logical expressions unlike Java where the case needs to be constants
		fmt.Println("first case")
	case 3, 6, 12:
		fmt.Println("second case")
	default:
		fmt.Println("Default case")
	}
	// If you want to execute multiple cases
	i := 5
	switch {
	case i <= 5:
		fmt.Println("Initial case")
		fallthrough // Possible to break out using break keyword
	case i >= 5: // This will be executed even if the condition is false
		fmt.Println("Fallthrough case")
	}
	// Checking data type using switch
	var arr interface{} = [...]int{1, 2, 3}
	switch arr.(type) {
	case int:
		fmt.Println("variable is integer")
	case [3]int: // need same size and type
		fmt.Println("Found array of integers")
	}

}
