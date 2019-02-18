package main


import (
	"fmt"
	"time"
	)


func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Data from Server1.."
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Data from Server2.."
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    output1 := make(chan string)
    output2 := make(chan string)

    go server1(output1)
    go server2(output2)

    time.Sleep(1 * time.Second)

    select {
    case s1 := <-output1 :
    	fmt.Println(s1)
    case s2 := <-output2 :
    	fmt.Println(s2)
    }


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run select.go 
Hello Go Developer.. :]
Data from Server2..
*/


/* => What is select?
-> The select statement is used to choose from multiple send/receive channel operations. The select 
	statement blocks until one of the send/receive operation is ready. If multiple operations are ready, 
	one of them is chosen at random. The syntax is similar to switch except that each of the case 
	statement will be a channel operation. Lets dive right into some code for better understanding.
*/
