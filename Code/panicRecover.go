package main


import (
	"fmt"
	)


func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from", r)
	}
}


func fullName(firstName string, lastName string) {
	defer recoverName()
	if firstName == "" {
		panic("Run Time err!, FirstName cann't be blank :(")
	}
	if lastName == "" {
		panic("Run Time err!, LastName cann't be blank :(")
	}
	fmt.Printf("FullName is = %s %s \n", firstName, lastName)
	fmt.Println("Returned normally from fullName() :)")
}



func main() {
    fmt.Println("Hello Go Developer.. :]")

    firstName := "Shamu"
    fullName(firstName, "")
    fmt.Println("Returned normally from main() :)")


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run panicRecover.go 
Hello Go Developer.. :]
Recovered from Run Time err!, LastName cann't be blank :(
Returned normally from main() :)																																						 
*/


/* => Recover
-> recover is a builtin function which is used to regain control of a panicking goroutine.
-> Recover is useful only when called inside deferred functions. Executing a call to recover 
	inside a deferred function stops the panicking sequence by restoring normal execution and 
	retrieves the error value passed to the call of panic. If recover is called outside the 
	deferred function, it will not stop a panicking sequence.
-> Recover works only when it is called from the same goroutine. It's not possible to recover 
	from a panic that has happened in a different goroutine.
*/