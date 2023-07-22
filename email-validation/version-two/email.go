package main

import (
	"fmt"
	"email-validation/check"
)

func main() {
	emails := []string{"emarjitkr@gamil.com", "amarjit@"}
	for index, email := range emails {
		if index == 1 {
			continue
		} else {
			if check.IsEmailValid(email){
				fmt.Println("Email is valid ", email)
			} else {
				fmt.Println("Email is not valid", email)
			}
		}
	}
}