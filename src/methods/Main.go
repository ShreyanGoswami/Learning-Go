package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	g.changeName("Shrey")
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

//Method
func (g greeter) greet() { // Copy of the struct is passed
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) changeName(name string) { // Pointer receiver
	g.name = name
}
