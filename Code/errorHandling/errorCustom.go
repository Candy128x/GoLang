package main


import (
	"fmt"
	"math"
	_"errors"
	)


// func circleArea(radius float64) (float64, error) {
// 	if radius < 0 {
// 		return 0, errors.New("Area calculation failed. Radius is less than zero.") // The simplest way to create a custom error is to use the New function of the errors package.
// 	}
// 	return math.Pi * radius * radius, nil
// }


// func circleArea2(radius float64) (float64, error) {
// 	if radius < 0 {
// 		return 0, fmt.Errorf("--2--Area calculation failed. Radius %0.2f is less than zero.", radius) // This is where the Errorf function of the fmt package comes in handy. This function formats the error according to a format specifier and returns a string as value that satisfies error.
// 	}
// 	return math.Pi * radius * radius, nil
// }


type areaError3 struct {
	err string
	radius float64
}

func (e *areaError3) Error() string { // Error() it's deterministic/static modification not allowed
 	return fmt.Sprintf("--3-2--Radius %0.2f: %s", e.radius, e.err)
}

func circleArea3(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError3{"--3-1--Radius is -negative", radius}
	}
	return math.Pi * radius * radius, nil
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    // radius := -20.0
    // area, err := circleArea(radius)
    // if err != nil {
    // 	fmt.Println(err)
    // 	return
    // }
    // fmt.Println("Area of circle =",area)

    /* Output:
	[error] if set radius := -20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	Area calculation failed. Radius is less than zero.

	[correct] if set radius := 20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	Area of circle = 1256.6370614359173
	*/


	/* => Adding more information to the error using Errorf.
	-> The above program works well but wouldn't it be nice if we print the actual radius which 
		caused the error. This is where the Errorf function of the fmt package comes in handy. This 
		function formats the error according to a format specifier and returns a string as value 
		that satisfies error.
	*/

    // radius2 := 20.0
    // area2, err2 := circleArea2(radius2)
    // if err2 != nil {
    // 	fmt.Println(err2)
    // 	return
    // }
    // fmt.Println("--2--Area of circle =",area2)

    /* Output:
	[error] if set radius2 := -20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	--2--Area calculation failed. Radius -20.00 is less than zero.

	[correct] if set radius2 := 20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	--2--Area of circle = 1256.6370614359173
	*/


	/* => This is where the Errorf function of the fmt package comes in handy. This function 
	formats the error according to a format specifier and returns a string as value that 
	satisfies error. */

	radius3 := -20.0
	area3, err3 := circleArea3(radius3)
	if err3 != nil {
		if err3, ok3 := err3.(*areaError3); ok3 {
			fmt.Printf("--3-3--Radius %0.2f is less than Zero. | Statment = %s \n", err3.radius, err3.err)
			return
		}
		fmt.Println("--3-4--", err3)
		return
	}
	fmt.Printf("--3-5--Area of Ractangle %0.2f", area3)

	/* Output:
	[error] if set radius3 := -20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	--3-3--Radius -20.00 is less than Zero. | Statment = --3-1--Radius is -negative 

	[correct] if set radius3 := 20.0
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
	Hello Go Developer.. :]
	--3-5--Area of Ractangle 1256.64
	*/


    fmt.Println("\n") // The End.. .
}


/* Output:
[error] if set radius3 := -20.0
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
Hello Go Developer.. :]
--3-3--Radius -20.00 is less than Zero. | Statment = --3-1--Radius is -negative 

[correct] if set radius3 := 20.0
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run errorCustom.go 
Hello Go Developer.. :]
--3-5--Area of Ractangle 1256.64

*/


/* Title
-> Description..
*/
