package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("The number %v is Even\n", i)
		} else {
			fmt.Printf("The number %v is Odd\n", i)
		}
	}

}
