package main


import (
	"fmt"
	"sync"
	)


var x = 0

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    // Solving the race condition using mutex

    var wg sync.WaitGroup
    var mu sync.Mutex

    for i := 0; i < 1000; i++ {
    	wg.Add(1)
    	go increment(&wg, &mu)
    }
    wg.Wait()

    fmt.Println("Final Value of x:", x)

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run mutex.go 
Hello Go Developer.. :]
Final Value of x: 1000
*/


/* => Mutex
-> A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the 
	critical section of code at any point of time to prevent race condition from happening.
*/
