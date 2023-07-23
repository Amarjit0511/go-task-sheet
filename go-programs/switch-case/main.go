package main

import "fmt"

func main() {
    var day int

    fmt.Print("Enter the day number (1-7): ")
    _, err := fmt.Scan(&day)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day number. Please provide a number between 1 and 6")
    }
}
