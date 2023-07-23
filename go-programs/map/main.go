package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	capitals := make(map[string] string) 

	capitals["Jharkhand"] = "Ranchi"
	capitals["Bihar"] = "Patna"
	capitals["West Bengal"] = "Kolkata"

	fmt.Println("List of capitals", capitals)
}
