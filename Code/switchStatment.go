package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    // fruit := "chikku"
    fruit := "green grapes"

    switch fruit {
    case "apple" :
    	fmt.Println("Your fav fruit is apple.")

	case "banana" :
		fmt.Println("Your fav fruit is banana.")
    
	case "chikku" :
		fmt.Println("Your fav fruit is chikku.")
	// case "chikku" :
	// 	fmt.Println("Your fav fruit is 3chikku.") // err: duplicate case "chikku" in switch

	case "green grapes", "black grapes" :
		fmt.Println("Your fav fruit is grapes.")

	default : 
		fmt.Println("Your fav fruit is non.")
    }


    fmt.Println("\n\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run switchStatment.go 
Hello Go Developer.. :]
Your fav fruit is grapes.

*/