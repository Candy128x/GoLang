package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")


    for i := 1; i <= 10; i++ {
    	fmt.Print(" ",i)
    }
    fmt.Println("\n")
    // op:1 2 3 4 5 6 7 8 9 10


    for i := 1; i <= 10; i+=2 {
    	fmt.Print(" ",i)
    }
    fmt.Println("\n")
    // op: 1 3 5 7 9


    for num := 0; num <= 10; num++ {  
	    if num % 2 == 0 {
	    	fmt.Println("Even number:", num)
	    } else {
	    	fmt.Println("Odd number;", num)
	    }
	}


	fmt.Println("\n break statment")
	/* The continue statement is used to skip the current iteration of the for loop. All code present in 
		a for loop after the continue statement will not be executed for the current iteration. The loop 
		will move on to the next iteration.
	*/
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // terminated.. . & exit
			fmt.Print(" ", i)
		}
		fmt.Print(" ", i)
	}
	// op: 1 2 3 4 5


	fmt.Println("\n\n break statment")
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            break
        }
        fmt.Print(" ", i)
    }
    /* Output:
     break statment
 	 1
    */

	fmt.Println("\n continue statment")
	for i := 1; i <= 10; i++ {
		if i %2 == 0 {
			continue // continue with next iteration..
			fmt.Print("--in if condition--", i)
		}
		fmt.Print(" ", i)
	}
    /* Output:
 	 continue statment
 	 1 3 5 7 9
    */


    fmt.Println("\n\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run forLoop.go 
Hello Go Developer.. :]
 1 2 3 4 5 6 7 8 9 10

 1 3 5 7 9

Even number: 0
Odd number; 1
Even number: 2
Odd number; 3
Even number: 4
Odd number; 5
Even number: 6
Odd number; 7
Even number: 8
Odd number; 9
Even number: 10

 break statment
 1 2 3 4 5

 break statment
 1
 continue statment
 1 3 5 7 9

*/