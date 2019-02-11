package main

import (
	"fmt"
	"time"
	)

func main() {
    fmt.Println("Hello Go Developer.. :]")

    go hello()
    // fmt.Println("time.Second:", time.Second)
    // fmt.Println("time.Millisecond:", time.Millisecond)
    time.Sleep(1 * time.Millisecond)
    fmt.Println("main() func..")


    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\n main terminated.. .")
}


func hello() {
	fmt.Println("Hello goRoutines..")
}


func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	} 
}


func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}


/*
why our Goroutine did not run. After the call to go hello() in line no. 11, the control returned 
	immediately to the next line of code without waiting for the hello goroutine to finish and printed 
	main function. Then the main Goroutine terminated since there is no other code to execute and hence 
	the hello Goroutine did not get a chance to run.
*/