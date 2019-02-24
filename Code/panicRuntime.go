package main


import (
	"fmt"
	)

	
func a() {
	defer r()
	n := []int {5, 7, 4}
	fmt.Println("Value of array n =", n[3])
	fmt.Println("Normally returned from a() :)")
}

func r() {
	if r := recover(); r != nil {
	fmt.Println("Recovered =", r)
	}
}
	

func main() {
    fmt.Println("Hello Go Developer.. :]")
	
	a()
	fmt.Println("Normally returned from main() :)")
	
    fmt.Println("\n") // The End.. .
}


/* Output:
[incorrect] i set -> fmt.Println("Value of array n =", n[3]) & without func a() -> defer r() & r()
E:\4-Data\GoLang\Concept\Code>go run panicRuntime.go
Hello Go Developer.. :]
panic: runtime error: index out of range

goroutine 1 [running]:
main.a()
        E:/4-Data/GoLang/Concept/Code/panicRuntime.go:11 +0x11
main.main()
        E:/4-Data/GoLang/Concept/Code/panicRuntime.go:19 +0x6d
exit status 2


--2--
[correct]
E:\4-Data\GoLang\Concept\Code>go run panicRuntime.go
Hello Go Developer.. :]
Recovered = runtime error: index out of range
Normally returned from main() :)
*/


/* => Runtime panics
-> Panics can also be caused by runtime errors such as array out of bounds access. This is 
	equivalent to a call of the built-in function panic with an argument defined by interface 
	type runtime.Error.
*/
