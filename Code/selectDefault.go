package main


import (
	"fmt"
	"time"
	)


func process(ch chan string) {
    time.Sleep(10500 * time.Millisecond)
    ch <- "Process Successful :)"
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan string)
    go process(ch)

    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("Recieved Value:", v)
            return
        default :
            fmt.Println("No Value Recieved!")
        }
    }


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run selectDefault.go 
Hello Go Developer.. :]
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
No Value Recieved!
Recieved Value: Process Successful :)
*/


/* => What is select default ?
-> The default case in a select statement is executed when none of the other case is ready. This is 
    generally used to prevent the select statement from blocking.
*/
