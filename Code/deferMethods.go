package main


import (
	"fmt"
	)


type person struct {
	firstName, lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s \n", p.firstName, p.lastName)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    p := person {"Ramu", "Dave"}
	defer p.fullName()
	fmt.Printf("Welcome ")
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run deferMethods.go 
Hello Go Developer.. :]
Welcome Ramu Dave 

*/


/* Title
-> Description..
*/
