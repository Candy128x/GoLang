package main

import "fmt"

func main() {
    fmt.Println("--Go Note </>--")

// *** Declare & initialize variables.. ***
    var a int = 10 // tightly couple
    a := 20 // loose couple

	// Data types
    var a int = 5
    var b float32 = 4.56
    const pi float64 = 3.1451
    x,y := 23, 34

    fmt.Println(a, b, pi, "\n", x, y)

	// Operators    
    fmt.Println("x / y = ", x/y)
    fmt.Println("x * y = ", x*y)
    fmt.Println("x + y = ", x+y)
    fmt.Println("x - y = ", x-y)
    fmt.Println("x mod y = ", x%y)


    isBool := true
    hate := false

    fmt.Println(isBool && hate)
    fmt.Println(isBool || hate)
    fmt.Println(!isBool)

 	// Pointers   
    fmt.Println(x)
    fmt.Println(&x)
    fmt.Println(&y)

    // changeValue(x) // x as value 
    changeValueByRef(&x) // x as reference
    fmt.Println(x)

}

func changeValue(x int){
	x = 13;
}
func changeValueByRef(x *int){
	*x = 17;
}


// *** Method decaration ***
func functionName(x int){
	x = x*x
}

    // for private use..
func loginUser(){	// fun name start with small caps  

}

	// for public use..
func ValidateUserPermission(){	// fun name start with Capital

}