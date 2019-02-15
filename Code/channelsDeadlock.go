package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    ch := make(chan string, 2)
    ch <- "Audi"
    ch <- "BMW"
    // ch <- "Chevrolet"

    close(ch)

    for v := range ch {
    	fmt.Println("Values of v:", v)
    }

    fmt.Println("\n") // The End.. .
}


/* Output:
=> if we uncommit -> ch <- "Chevrolet"
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsDeadlock.go 
Hello Go Developer.. :]
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/ashishs/Downloads/storage/Go/GoLang/Code/channelsDeadlock.go:15 +0x16a
exit status 2


=> if we commit -> // ch <- "Chevrolet"
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channelsDeadlock.go 
Hello Go Developer.. :]
Values of v: Audi
Values of v: BMW
*/


/* Title
-> Description..
*/
