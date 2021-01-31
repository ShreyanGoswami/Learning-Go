package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Error: cannot divide by 0")
		}
		return a / b, nil
	}
	d, _ := divide(5.0, 3.0)
	fmt.Println(d)
}
