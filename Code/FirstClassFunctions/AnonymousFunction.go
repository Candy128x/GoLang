package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

	/* ## Anonymous function */
	a := func() {
		fmt.Println("Hey Dude ;) from First Class Function.")
	}
	a()
	fmt.Printf("type = %T", a)
	
	/* ### These kind of functions are called anonymous functions since they do not have a name. */
	
	
	func(){
		fmt.Println("\n\nI dont have a func name !")
	}()
	
	/* ### It is also possible to call a anonymous function without assigning it to a variable. */

	
	func(n string) {
		fmt.Println("\nHello", n, ":]")
	}("Gopher")
	
	/* ### It is also possible to pass arguments to anonymous functions just like any other function. */
	
    fmt.Println("\n") // The End.. .
}


/* Output:
E:\4-Data\GoLang\Concept\Code\FirstClassFunctions>go run AnonymousFunction.go
Hello Go Developer.. :]
Hey Dude ;) from First Class Function.
type = func()

I dont have a func name !

Hello Gopher :]
*/


/* => What are first class functions?
-> A language which supports first class functions allows functions to be assigned to variables, 
	passed as arguments to other functions and returned from other functions. Go has support for 
	first class functions.
*/
