package main


import (
	"fmt"
	)

	
func appendStr() func(string) string {  
    t := "Hello"
    c := func(b string) string {
        t = t + " " + b
        return t
    }
    return c
}
	

func main() {
    fmt.Println("Hello Go Developer.. :]")

	a1 := 5
	func() {
		fmt.Println("a1 =", a1)
	}()

	
	/* --2-- */
	a := appendStr()
    b := appendStr()
    fmt.Println(a("World"))
    fmt.Println(b("Everyone"))

    fmt.Println(a("Gopher"))
    fmt.Println(b("!"))
    fmt.Println("\n") // The End.. .
}


/* Output:
E:\4-Data\GoLang\Concept\Code\FirstClassFunctions>go run Closures.go
Hello Go Developer.. :]
a1 = 5
Hello World
Hello Everyone
Hello World Gopher
Hello Everyone !
*/


/* Closures
-> Closures are a special case of anonymous functions. Closures are anonymous functions which access 
	the variables defined outside the body of the function. 
->  the anonymous function accesses the variable a which is present outside its body in code >>> fmt.Println("a =", a) <<<. Hence this anonymous function is a closure.
*/
