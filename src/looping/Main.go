package main

import "fmt"

func main() {
	// Simple loop
	fmt.Println("Simple loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// Loop with two variables
	fmt.Println("For loop with two variables")
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	// Do while loops are not present by default but you can change
	// the for loop to act as a do while
	fmt.Println("Do While loop")
	doWhileCounter := 0
	for doWhileCounter < 5 {
		fmt.Println(doWhileCounter)
		doWhileCounter++
	}
	// Breaking out of nested loops using labels
	fmt.Println("Breaking out of nested loops")
Loop:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	// Working with collections in loops(for range loop)
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Printf("Index %v Value %v\n", k, v)
	}

}
