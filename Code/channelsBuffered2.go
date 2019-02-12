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

	// close(ch)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan int, 1)
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
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run Demo.go 
Hello Go Developer.. :]
*/


/* Desc:
-> In this example of buffered channel in which the values to the channel are written in a concurrent 
	Goroutine and read from the main Goroutine. This example will help us better understand when 
	writes to a buffered channel block.
*/
