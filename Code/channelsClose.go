package main


import (
	"fmt"
	)


func producer(chnl chan int) {
	for i := 1; i < 10; i++ {
		chnl <- i // write i value into chnl
	}
	close(chnl) // closes the channels
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan int)
    go producer(ch)
    for {
    	v, ok := <-ch
    	if ok == false {
    		break
    	}
    	fmt.Println("Received data from chnl:", v, ok)
    }


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsClose.go 
Hello Go Developer.. :]
Received data from chnl: 1 true
Received data from chnl: 2 true
Received data from chnl: 3 true
Received data from chnl: 4 true
Received data from chnl: 5 true
Received data from chnl: 6 true
Received data from chnl: 7 true
Received data from chnl: 8 true
Received data from chnl: 9 true
*/


/* Title
-> Description..
*/
