package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    fruit := "AppleFruit"
    fmt.Println("Orignal string:", fruit)

    fmt.Println("Reverse string...")
    for _, v := range []rune(fruit) {
    	defer fmt.Printf("%c", v)
    }
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run deferStack.go 
Hello Go Developer.. :]
Orignal string: AppleFruit
Reverse string...
tiurFelppA
*/


/* Title
-> Description..
*/
