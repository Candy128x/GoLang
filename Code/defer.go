package main


import (
	"fmt"
	)

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func finished() {
	fmt.Println("Finished finding largest.")
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    nums := []int{23, 34, 54, 65, 73, 44, 56}
    largest(nums)

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run defer.go 
Hello Go Developer.. :]
Started finding largest
Largest number in [23 34 54 65 73 44 56] is 73
Finished finding largest.
*/


/* => Defer?
-> The statement defer finished(). This means that the finished() function will be called just before 
	the largest function returns. Run this program and you can see the following output printed.
-> It's trigger when function executions are terminated.
-> We cant wrote multiple defer statment in single func.
*/
