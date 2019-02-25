package main


import (
	"fmt"
	"reflect"
	)


type order struct {
	orderId, customerId int
}


func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	fmt.Println("\n\n Reflect, TypeOf q =", t)
	v := reflect.ValueOf(q)
	fmt.Println("\n Reflect, ValueOf q =", v)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    i := 10
    fmt.Printf("Value of i = %d | Type = %T", i, i)

    o := order {
    	384,
    	23,
    }
    createQuery(o)

    fmt.Println("\n") // The End.. .
}



/* Output:
[ACTUAL]
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/Reflection$ go run main.go 
Hello Go Developer.. :]
Value of i = 10 | Type = int

 Reflect, TypeOf q = main.order

 Reflect, ValueOf q = <main.order Value>

[EXCEPTED]
@ref: https://golangbot.com/reflection/
Type  main.order
Value  {456 56}
*/