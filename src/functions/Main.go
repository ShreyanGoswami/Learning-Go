package main

import "fmt"

func main() {
	sayMessage("Hello Go!")
	sayGreeting("Hi", "Shrey")
	fmt.Println(sum([]int{1, 2, 3}...))
	fmt.Println(sumPtr([]int{1, 2, 3, 4, 5}...))
	fmt.Println(sumAutoReturn([]int{1, 2, 3, 4, 5}...))
	result, err := safeDivide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func sayMessage(msg string) {
	fmt.Printf("Message: %v\n", msg)
}

func sayGreeting(greeting, name string) { // Compiler infers that greeting and name have same type
	fmt.Println(greeting, name)
}

// variadic parameter
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sumPtr(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // Unsafe operations in other language but Go will automatically promote this variable into the heap rather than having it on the stack and gc later
}

func sumAutoReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}
