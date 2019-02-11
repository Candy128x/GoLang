package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    age := 23
    if 18 <= age {
    	fmt.Println("Your age are grater than 18")
    }

    food := "veg"
    if "nonVeg" == food {
    	fmt.Println("You are nonVeg :(")
    } else {
    	fmt.Println("You are veg :)")
    }

    if num := 10; num % 2 == 0 {
    	fmt.Println("Even number:", num)
    } else {
    	fmt.Println("Odd number;", num)
    }

    num := 99 
    if num <= 50 {
    	fmt.Println("Number are less than 50..")
    } else if num >= 51 && num <=100 {
    	fmt.Println("Number between 51 to 100..")
    } else {
    	fmt.Println("Number are grater than 101..")
    }

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run ifelse.go 
Hello Go Developer.. :]
Your age are grater than 18
You are veg :)
Even number: 10
Number between 51 to 100..

*/