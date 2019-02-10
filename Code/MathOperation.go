package main

import (
	"fmt"
	"math"
)

func main() {

	a, b := 12.25, 12.75
	min := math.Min(a, b)
	fmt.Println("Minimum value is:", min)
	
	/* Output:
	E:\4-Data\GoLang\Concept\Code>go run MathOperation.go
	Minimum value is: 12.25
	*/
}