package main

import "fmt"

const (
	_ = iota // Throw away variable, won't be used because there is no need for 0 value
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

// constant block for size
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// Constants - can't be used in arrays and can be shadowed
	// Can be type inferred
	const myConst int = 42 // We don't want to start with upper case because then it will expose the constant
	fmt.Printf("%v, %T\n", myConst, myConst)
	// Enumerated constants
	const (
		a = iota
		b = iota
		c = iota
	)
	/*  Better way
	const (
		a = iota
		b
		c
	)
	*/
	fmt.Printf("%v\n", a) // 0
	fmt.Printf("%v\n", b) // 1
	fmt.Printf("%v\n", c) // 2
	// Convert to GB
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
