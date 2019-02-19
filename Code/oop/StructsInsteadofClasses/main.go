package main


import (
	"fmt"
	"oop/employee"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    // Structs Instead of Classes - OOP in Go..
    e := employee.New("Shamu", "Jain", 30, 20)
    e.LeavesRemaining()

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/oop/StructsInsteadofClasses$ go run main.go 
Hello Go Developer.. :]
Shamu Jain has 10 leaves remaining.
*/


/* Title
-> Description..
*/
