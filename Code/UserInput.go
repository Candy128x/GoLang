package main

import "fmt"

func main() {
    fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)
    fmt.Print("Have a Good Day " + input + "!")
    fmt.Print("\n How are you ", input, "..")
}

/* Output:
E:\4-Data\GoLang\Concept\Code>go run UserInput.go
Enter text: Ashish
Have a Good Day Ashish!
 How are you Ashish..
*/