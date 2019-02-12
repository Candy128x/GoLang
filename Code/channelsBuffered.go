package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan string, 2)
    ch <- "Ramu" // write 1st
    ch <- "Shamu" // write 2nd
    // ch <- "kaku" // write 3rd | get compilation err if uncommit
    
    fmt.Println("ch data:",<- ch) // reads 1st
    fmt.Println("ch data:",<- ch) // reads 2nd
    
    // fmt.Println("ch data:",<- ch) // reads 3rd ? | get compilation err if uncommit
    //get err :( => fatal error: all goroutines are asleep - deadlock!


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsBuffered.go 
Hello Go Developer.. :]
ch data: Ramu
ch data: Shamu
*/


/* Title
-> Description..
*/
