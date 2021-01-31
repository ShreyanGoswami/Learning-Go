package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)
	// Declaring floats - follows IEEE754, 32 bit and 64 bit
	x := 3.14
	x = 13.7e72
	x = 2.1e14
	fmt.Printf("%v, %T\n", x, x)
	// Complex numbers
	var c complex64 = 1 + 2i // float32 + float32
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("Real part %v Complex part %v \n", real(c), imag(c))
	// Create complex numbers
	var custom complex64 = complex(5, 12)
	fmt.Printf("%v %T\n", custom, custom)
	// Strings
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
	// Runes(UTF-32)
	r := 'a'
	var explicitRune rune = 'b'
	fmt.Printf("%v, %T\n", r, r)
	fmt.Printf("%v, %T\n", explicitRune, explicitRune)
	// There are special functions in the go language to read runes such as ReadRune
}
