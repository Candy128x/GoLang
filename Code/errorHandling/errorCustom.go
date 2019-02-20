package main


import (
	"fmt"
	"math"
	"errors"
	)


func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed. Radius is less than zero.") // The simplest way to create a custom error is to use the New function of the errors package.
	}
	return math.Pi * radius * radius, nil
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    fmt.Println("Area of circle =",area)


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
Hello Go Developer.. :]
Area calculation failed. Radius is less than zero.
*/


/* Title
-> Description..
*/
