package main


import (
	"fmt"
	)


type add func(no1, no2 int) int 
	
	
func main() {
    fmt.Println("Hello Go Developer.. :]")

	/* ## User defined function types */

	var a add = func(no1, no2 int) int {
		return no1 + no2
	}
	
	s := a(5,6)
	fmt.Println("Result of s =", s)
	
    fmt.Println("\n") // The End.. .
}


/* Output:
E:\4-Data\GoLang\Concept\Code\FirstClassFunctions>go run UserDefinedFunction.go
Hello Go Developer.. :]
Result of s = 11
*/


/* User defined function types
-> Just like we define our own struct types.
*/
