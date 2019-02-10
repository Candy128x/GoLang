package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calculate Bill:", calculateBill(500,3))
	
	area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f", area, perimeter)
	
	area2, perimeter2 := rectProps2(2.8, 3.6)
    fmt.Printf("\n Area %f Perimeter %f", area2, perimeter2)
	
	// blank identifier
	area3, _ := rectProps2(4.8, 6.6)
    fmt.Printf("\n Area %f", area3)

}

func calculateBill(price, no int) int {
	return price * no
}


// Multiple return values
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}


// Named return values
func rectProps2(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}


/*
-> What is a function?
- A function is a block of code that performs a specific task. A function takes a input, performs some 
	calculations on the input and generates a output.
- Syntax: 
	func functionname(parametername type) returntype {  
	//function body
	}

*/