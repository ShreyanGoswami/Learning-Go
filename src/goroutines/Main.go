package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// Spawn a goroutine but nothing will happen if we don't wait
	go sayHello()
	// With anonymous functions
	msg := "Hello"
	go func() {
		fmt.Println(msg) // Go implements closure so the go routine knows that the variable defined outside is required inside. Its not a good practice.
	}()
	time.Sleep(100 * time.Millisecond)
	// Better way
	safeMsg := "Hello"
	go func(msg string) { // Copied by value
		fmt.Println(msg)
	}(safeMsg)
	time.Sleep(100 * time.Millisecond)
	// How to avoid sleep? With wait group
	waitGroupMsg := "Running in wait group"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(waitGroupMsg)
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
}
