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
	
	k := t.Kind()
	fmt.Println("\n Reflect, Kind =", k)

	nf := v.NumField()
	fmt.Println("\n Reflect, NumField =", nf)

	for i := 0; i < nf; i++ {
		fmt.Printf("\n Reflect, Field: %d | Type: %T | Value:%v \n", i, v.Field(i), v.Field(i))
	}

	int1 := 23
	vInt := reflect.ValueOf(int1).Int()
	fmt.Printf("\n Reflect, Value of int1 = %v | Type = %T", vInt, vInt)

	str1 := "Gopher"
	vStr := reflect.ValueOf(str1).String()
	fmt.Printf("\n Reflect, Value of str1 = %v | Type = %T", vStr, vStr)
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
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/Reflection$ go run main.go 
Hello Go Developer.. :]
Value of i = 10 | Type = int

 Reflect, TypeOf q = main.order

 Reflect, ValueOf q = <main.order Value>

 Reflect, Kind = struct

 Reflect, NumField = 2

 Reflect, Field: 0 | Type: reflect.Value | Value:<int Value> 

 Reflect, Field: 1 | Type: reflect.Value | Value:<int Value> 

 Reflect, Value of int1 = 23 | Type = int64
 Reflect, Value of str1 = Gopher | Type = string
*/


/* => What is reflection?
-> Reflection is the ability of a program to inspect its variables and values at run time and find 
	their type. You will get a clear 
	understanding of reflection by the end of this tutorial, so stay with me.
-> reflect package
	reflect.Type and reflect.Value
	reflect.Kind
	NumField() and Field() methods
	Int() and String() methods

=> reflect package
-> The reflect package implements run-time reflection in Go. The reflect package helps to identify 
	the underlying concrete type and the value of a interface{} variable. This is exactly what 
	we need. The createQuery function takes a interface{} argument and the query needs to be created 
	based on the concrete type and value of the interface{} argument. This is exactly what the 
	reflect package helps in doing.
	1. reflect.Type and reflect.Value
	2. reflect.Kind
		- There is one more important type in the reflection package called Kind.
		- The types Kind and Type in the reflection package might seem similar but they have a 
			difference which will be clear from the program below.
	3. NumField() and Field() methods
	4. Int() and String() methods
		- The methods Int and String help extract the reflect.Value as an int64 and string respectively.
*/
