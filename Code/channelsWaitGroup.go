package main


import (
	"fmt"
	"sync"
	"time"
	)


func process(i int, wg *sync.WaitGroup){
	fmt.Println("Started goRoutines", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("GoRoutines %d ended \n", i)
	wg.Done()
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    no := 3

    var wg sync.WaitGroup
    // WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup

    for i := 0; i < no; i++ {
    	wg.Add(1)
    	go process(i, &wg)
    }
    
    wg.Wait()
    /*  
    	When we call Add on the WaitGroup and pass it an int, the WaitGroup's counter is incremented 
    by the value passed to Add. The way to decrement the counter is by calling Done() method on the 
    WaitGroup. The Wait() method blocks the Goroutine in which it's called until the counter becomes zero.
    */

    fmt.Println("All goRoutines finish executing.. .");


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run Demo.go 
Hello Go Developer.. :]
*/


/*
=>Worker Pools. 
-> To understand worker pools, we need to first know about WaitGroup as it will be used in the 
	implementation of Worker pool.
-> A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is 
	blocked until all Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines 
	spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to 
	finish before terminating. This can be accomplished using WaitGroup.
*/
	