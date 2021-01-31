package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	// Dereference
	fmt.Println(a, *b)
	// Changing the value of a from b
	*b = 24
	fmt.Println(a, *b)
	// Pointers and structs
	var ms *myStruct = &myStruct{foo: 42}
	fmt.Println(ms)
	// Using new to create a pointer to a struct
	var ptr *myStruct = new(myStruct)
	(*ptr).foo = 24
	fmt.Println(ptr)
	// OR a better way
	ptr.foo = 42 // compiler optimisation
	fmt.Println(ptr)
	// Slices and maps are pointers. Primitives, arrays and structs are copies unless explicitly specified
}
