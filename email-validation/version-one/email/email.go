package main
import (
    "fmt"
    "email-validation/check"
)
func main() {
    // List of emails to check
    emails := []string{"amarjitkr0511@gmail.com", "amarjitkr@.com", "amarjit@kr@gmail.com"}

    // Looping through each email
    for _, email := range emails {
        // Checking if the email is valid or invalid
        if check.IsEmailValid(email) {
            fmt.Printf("%s is a valid email.\n", email)
        } else {
            fmt.Printf("%s is an invalid email.\n", email)
        }
    }
}
