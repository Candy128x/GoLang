package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    a := 10
    fmt.Printf("value of a:%d | type %T", a, a)
    fmt.Println("\naddrs value", &a)

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run Demo.go 
Hello Go Developer.. :]
*/