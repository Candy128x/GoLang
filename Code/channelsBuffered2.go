package main


import (
	"fmt"
	"time"
	)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote:", i ,"to ch :)")
	}
	close(ch)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)

    for v := range ch {
    	fmt.Println("Read value:", v,"from ch")
    	// fmt.Println("Length of ch:", len(ch))
    	time.Sleep(2 * time.Second)
    }

    fmt.Println("Length of ch:", len(ch))
    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsBuffered2.go
Hello Go Developer.. :]
Successfully wrote: 0 to ch :)
Successfully wrote: 1 to ch :)
Read value: 0 from ch
Successfully wrote: 2 to ch :)
Read value: 1 from ch
Successfully wrote: 3 to ch :)
Read value: 2 from ch
Successfully wrote: 4 to ch :)
Read value: 3 from ch
Read value: 4 from ch
Length of ch: 0
*/


/* Desc:
-> In this example of buffered channel in which the values to the channel are written in a concurrent 
	Goroutine and read from the main Goroutine. This example will help us better understand when 
	writes to a buffered channel block.

@ref: https://golangbot.com/buffered-channels-worker-pools/
*/
