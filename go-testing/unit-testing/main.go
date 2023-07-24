package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func division(a, b int) int {
	if b == 0 {
		panic("Not valid")
	}
	return a / b
}

func multiplication(a, b int) int {
	return a * b
}

func main() {
	a, b := 10, 5

	fmt.Printf("Addition: %d + %d = %d\n", a, b, addition(a, b))
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, subtraction(a, b))
	fmt.Printf("Division: %d / %d = %d\n", a, b, division(a, b))
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, multiplication(a, b))
}
