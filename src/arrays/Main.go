package main

import "fmt"

func main() {
	// declare array
	grades := [3]int{97, 85, 93} // contiguous in memory
	// grades := [...]int{97, 85, 93}
	fmt.Printf("Grades %v\n", grades)
	fmt.Printf("Size of array: %v\n", len(grades))

	// 2D array
	var identityMatrix [3][3]int
	identityMatrix[0] = [...]int{1, 0, 0}
	identityMatrix[1] = [...]int{0, 1, 0}
	identityMatrix[2] = [...]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// Arrays are copies
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// Slices(works by reference not copies)
	slice := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n", cap(slice))

	// Make
	newSlice := make([]int, 3, 100)
	fmt.Println(newSlice)

	// Appending
	variableSizeSlice := []int{}
	variableSizeSlice = append(variableSizeSlice, 1) // be careful with this as it might trigger an expensive copy operation if the underlying array is too small. You can use the 3 parameter make function to have an array larger than the length
	fmt.Printf("Length: %v\n", len(variableSizeSlice))
	fmt.Printf("Capacity: %v\n", cap(variableSizeSlice))

	// Concatenate
	initial := []int{1, 2, 3}
	initial = append(initial, []int{4, 5, 6}...) // Spread operator
	fmt.Printf("Concatenate result: %v\n", initial)
	fmt.Printf("Capacity: %v\n", cap(initial))

}
