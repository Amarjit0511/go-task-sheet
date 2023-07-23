package main

import "fmt"

func main() {
	fmt.Println("Slices in Golang")

	cities := []string{"Dumka", "Ranchi", "Lucknow", "Shimla"}
	fmt.Println(cities)

	cities = append(cities, "Kolkata")
	fmt.Println(cities)
}
