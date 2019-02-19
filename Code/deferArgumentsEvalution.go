package main


import (
	"fmt"
	)

func PrintA(a int) {
	fmt.Println("Value of a on defer call =", a)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    a := 5
    defer PrintA(a)
    a = 10
    fmt.Println("Value of a before defer call =", a)
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run deferArgumentsEvalution.go 
Hello Go Developer.. :]
Value of a before defer call = 10
Value of a on defer call = 5

*/


/* Title
-> Description..
*/
