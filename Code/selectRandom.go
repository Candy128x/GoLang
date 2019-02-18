package main


import (
	"fmt"
	"time"
	)


func server1(ch chan string) {
	ch <- "Data from Server1.."
}

func server2(ch chan string) {
	ch <- "Data from server2.."
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
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run selectRandom.go 
Hello Go Developer.. :]
Data from Server1..


ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run selectRandom.go 
Hello Go Developer.. :]
Data from server2..


ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run selectRandom.go 
Hello Go Developer.. :]
Data from server2..

@note
-> Every time we get diffrent output..
*/


/* Random Selection
-> When multiple cases in a select statement are ready, one of them will be executed at random.
*/
