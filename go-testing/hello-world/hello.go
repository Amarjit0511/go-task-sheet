package main

import "fmt"

func HelloWorld() string {
	return "Hello World"
}
func main() {
	message := HelloWorld()
	fmt.Println(message)
}
