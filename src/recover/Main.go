package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
	globalPanic()
	fmt.Println("This wont execute")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func globalPanic() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
