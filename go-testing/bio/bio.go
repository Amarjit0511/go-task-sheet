package main

import (
	"fmt"
	"strconv"
)

func Bio(name string, age int, gender string) []string {
	return []string{
		name,
		strconv.Itoa(age),
		gender,
	}
}

func main() {
	message := Bio("Amarji", 23, "Male")
	fmt.Println(message)
}
