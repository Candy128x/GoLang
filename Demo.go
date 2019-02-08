package main

import "fmt"

func main() {
    fmt.Println("Hello Go Developer.. :]")

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

}
