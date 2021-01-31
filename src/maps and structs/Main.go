package main

import (
	"fmt"
	"reflect"
)

type animal struct {
	id   int
	name string `required max:"100"` // Tag can be used for validation like a comment
}

type bird struct {
	animal // Embedding
	origin string
}

func main() {
	// maps are not in order and by reference
	foodTypeRating := map[string]int{ // Keys should be tested for equality such as string, boolean. Slices and maps cannot be tested for equality
		"Italian":  5,
		"Japanese": 5,
		"Indian":   4,
	}
	fmt.Println(foodTypeRating)

	// If you don't have the entries at the time of declaring
	myMap := make(map[string]int)
	myMap = map[string]int{
		"val1": 1,
		"val2": 2,
	}
	fmt.Println(myMap)

	// Map manipulation
	myMap["val3"] = 3
	fmt.Println(myMap["val3"])
	delete(myMap, "val3")
	missingVal, ok := myMap["val3"] // comma okay syntax
	fmt.Println(missingVal, ok)

	// Struct - works by copies
	aDog := animal{
		id:   1,
		name: "Dog",
	}
	// Accessing tags
	t := reflect.TypeOf(animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(aDog)
	fmt.Println(aDog.id)

	// Anonymous struct for short term usage
	randomStruct := struct{ name string }{name: "Shreyan"}
	fmt.Println(randomStruct)

	// using embedding with bird
	b := bird{}
	b.id = 2
	b.name = "Emu"
	b.origin = "Australia"
	fmt.Println(b)
}
