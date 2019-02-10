package main

import (
	"fmt"
)

func main() {
	const a = 23
	fmt.Println("Value of const a :", a)
	
	const strHello = "Hello Go Lang.."
	const strHelloCode string = "Developer"
	fmt.Println(strHello, strHelloCode)
	// op: Hello Go Lang.. Developer
	
	var name = "Sam"
    fmt.Printf("type %T value %v", name, name)
	
	
	var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed
	// customName = defaultName //not allowed
	fmt.Println("\n", defaultName, customName)
	
}

/* Output:
E:\4-Data\GoLang\Concept\Code>go run Constants.go
Value of const a : 23
Hello Go Lang.. Developer
type string value Sam
 Sam Sam
*/