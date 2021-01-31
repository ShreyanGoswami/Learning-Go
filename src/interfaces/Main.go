package main

import "fmt"

/*
Don't export interfaces for types that will be consumed
Do export interfaces for types that will be used by the package
*/

func main() {
	// Simple writer interface
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
	// Simple increment interface
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	return fmt.Println(string(data))
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
