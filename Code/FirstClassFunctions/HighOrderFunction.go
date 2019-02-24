package main


import (
	"fmt"
	)


func simple(a func(a, b int) int) {
	fmt.Println(a(70, 3))
}


func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

	f := func(a, b int) int {
		return a + b
	}
	simple(f)
	
	
	/* Returning functions from other functions */
	s2 := simple2()
	fmt.Println(s2(60, 4))
	
    fmt.Println("\n") // The End.. .
}


/* Output:
E:\4-Data\GoLang\Concept\Code\FirstClassFunctions>go run HighOrderFunction.go
Hello Go Developer.. :]
73
64
*/


/* => Higher-order functions
-> The definition of Higher-order function from wiki is a function which does at least one of the following
- 1 takes one or more functions as arguments
- 2 returns a function as its result
*/
