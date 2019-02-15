package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan string, 3)
    ch <- "Audi"
    ch <- "BMW"

    fmt.Println("Length of ch:", len(ch))
    fmt.Println("Capacity of ch:", cap(ch))
    fmt.Println("Read value from ch:", <-ch)
    fmt.Println("new Length of ch:", len(ch))


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsLengthCapacity.go 
Hello Go Developer.. :]
Length of ch: 2
Capacity of ch: 3
Read value from ch: Audi
new Length of ch: 1
*/


/* Title
-> Description..
*/
